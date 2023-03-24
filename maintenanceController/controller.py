from flask import Flask, request, jsonify
import os
from dotenv import load_dotenv, environ
import requests

import amqp_setup
import pika

app = Flask(__name__)

################# API Routes ######################
maintenanceAPI = 'http://127.0.0.1:5000/maintenance'
equipmentAPI = 'http://127.0.0.1:4999/equipment'
inventoryAPI = 'http://127.0.0.1:5001/inventory'



@app.route("/schedule_maintenance")
def scheduleMaintenance():
    '''
    1. Reserve Parts (HTTP)
    2. Create Maintenance Record
    3. Order Parts
    4. Send Notification
    '''
    data = request.get_json()
    eqp_id = data["equipment_id"]
    scheduled_datetime = data["schedule_datetime"]
    requested_parts = data["partlist"]

    # 1. Reserve Parts (HTTP)
    result = reserveParts(requested_parts)

    if type(result) != str:

        reserved_list, missing_list = result

        # 2. Create Maintenance Record
        data['partList'] = reserved_list

        maintenance_result = requests.request('POST', maintenanceAPI, json = data)
        maintenance_code = maintenance_result['code']

        amqp_setup.check_setup()

        # Successfully scheduled maintenance
        if maintenance_code == '201':

            # 3. Order Parts
            if len(missing_list):
                orderParts(missing_list)

            # 4. Send out Notification
            amqp_setup.channel.basic_publish(
                exchange=amqp_setup.exchangename, 
                routing_key="schedule.maintenance", 
                body=data, 
                properties=pika.BasicProperties(delivery_mode = 2)
            )


            return jsonify({
                "code": maintenance_code,
                "message": f"Maintenance for {eqp_id} has been scheduled on {scheduled_datetime}"
            })
        
        # Maintenance record already exist
        elif maintenance_code == '404':
            
            # Return reserved parts (RabbitMQ)
            returnParts(reserved_list)

            return jsonify({
                "code": maintenance_code,
                "message": "Maintenance record already exists"
            }),maintenance_code
        
        else:
            return jsonify({
                "code": maintenance_code,
                "message": maintenance_result['message']
            }),maintenance_code




    
@app.route("/start_maintenance/<string:maintenance_id>")
def startMaintenance(maintenance_id):
    '''
    1. Send update equipment (Status to maintenance)
    2. Send update maintenance (Start time, techID, Status)
    3. Send out Notification (Email + SMS)
    '''
    data = request.get_json()
    eqp_id = data["equipment_id"]
    # TechID
    eqp_status = "Maintenance In Progress"
    start_datetime = data["start_datetime"]
    maintenance_status = "STARTED"

    eqp_data = {
        "equipment_status": eqp_status
    }

    maintenance_data = {
        "start_datetime": start_datetime,
        "status": maintenance_status
    }

    eqp_result = requests.request('PUT', f'{equipmentAPI}/{eqp_id}', json = eqp_data)

    maintenance_result = requests.request('PUT', f'{maintenanceAPI}/{maintenance_id}', json = maintenance_data)
    maintenance_code = maintenance_result['code']

    if maintenance_code == '201':
        executeMaintenance(data, start=True)

        return jsonify({
            "code": maintenance_code,
            "message": "Maintenance successfully started"
        }),maintenance_code
    
    else:
        return jsonify({
            "code": maintenance_code,
            "message": maintenance_result['message']
        }),maintenance_code




@app.route("/request_parts/<string:maintenance_id>")
def requestParts(maintenance_id):
    '''
    1. Reserve Parts (HTTP)
    2. Send update maintenance (Partlist)
    '''
    

    pass



@app.route("/end_maintenance/<string:maintenance_id>")
def endMaintenance(maintenance_id):
    '''
    1. Send update maintenance (Last Service time, Status, Description)
    2. Send update equipment (Status to maintenance)
    3. Return Parts (RabbitMQ)
    3. Send out Notification (Email + SMS)
    '''
    data = request.get_json()
    eqp_id = data["equipment_id"]
    end_datetime = data["end_datetime"]
    maintenance_status = data["maintenance_status"]
    description = data["description"]
    part_list = data["partlist"]
    return_parts = data["return_partlist"]
    eqp_status = "No maintenance required" if maintenance_status == "COMPLETE - SUCCESSFUL" else "Maintenance required"


    return_parts_dict = {}
    for part in return_parts:
        return_parts_dict[part['_id']] = part['Qty']


    # Tabulate parts used
    for part in part_list:
        if part['_id'] in return_parts_dict:
            part['Qty'] = part['Qty'] - return_parts_dict[part['_id']]

    eqp_data = {
        "equipment_status": eqp_status
    }

    maintenance_data = {
        "end_datetime": end_datetime,
        "status": maintenance_status,
        "description": description,
        "partlist": part_list
    }

    maintenance_result = requests.request('PUT', f'{maintenanceAPI}/{maintenance_id}', json = maintenance_data)
    maintenance_code = maintenance_result['code']

    eqp_result = requests.request('PUT', f'{equipmentAPI}/{eqp_id}', json = eqp_data)

    if maintenance_code == '201':
        executeMaintenance(data, start=False)

        returnParts(return_parts)

        return jsonify({
            "code": maintenance_code,
            "message": "Maintenance successfully ended"
        }),maintenance_code
    
    else:
        return jsonify({
            "code": maintenance_code,
            "message": maintenance_result['message']
        }),maintenance_code




def reserveParts(partList):
    parts_result = requests.request('PUT', inventoryAPI, json=partList)
    code = parts_result["code"] 

    # Parts sucessfully reserved
    if code == '201':
        reservedList = parts_result['reservedList']
        missingList = parts_result['missingList']

        return (reservedList, missingList)
    
    else:
        return code['message']


def returnParts(partList):
    amqp_setup.channel.basic_publish(
        exchange=amqp_setup.exchangename, 
        routing_key="return.parts", 
        body=partList, 
        properties=pika.BasicProperties(delivery_mode = 2)
    )


def orderParts(partList):
    amqp_setup.channel.basic_publish(
        exchange=amqp_setup.exchangename, 
        routing_key="order.parts", 
        body=partList, 
        properties=pika.BasicProperties(delivery_mode = 2)
    )


def executeMaintenance(maintenanceData, start):
    if start:
        routing_key = "maintenance.start"
    else:
        routing_key = "maintenance.end"

    amqp_setup.channel.basic_publish(
        exchange=amqp_setup.exchangename, 
        routing_key=routing_key, 
        body=maintenanceData, 
        properties=pika.BasicProperties(delivery_mode = 2)
    )
