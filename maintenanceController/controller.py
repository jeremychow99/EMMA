from flask import Flask, request, jsonify
from flask_cors import CORS
# import os
# from dotenv import load_dotenv
import requests
import amqp_setup
import pika
import json

app = Flask(__name__)
CORS(app)

################# API Routes ######################
maintenanceAPI = 'http://127.0.0.1:5000/maintenance'
equipmentAPI = 'http://127.0.0.1:4999/equipment'
inventoryAPI = 'http://127.0.0.1:5001/inventory'



@app.route("/schedule_maintenance", methods=['POST'])
def scheduleMaintenance():
    '''
    1. Reserve Parts (HTTP)
    2. Create Maintenance Record
    3. Order Parts
    4. Send Notification
    '''
    print("Schedule Maintenance Function Triggered!")

    data = request.get_json()
    print(data)
    eqp_id = data["equipment"]["equipment_id"]
    schedule_date = data["schedule_date"]
    requested_parts = data["partlist"]

    # 1. Reserve Parts (HTTP)
    print("List of parts to reserve: ", requested_parts)
    print("Attempting to reserve parts...")
    result = reserveParts(requested_parts)

    if type(result) != str:
        print("Parts reserved...")
        reserved_list, missing_list = result

        # 2. Create Maintenance Record
        data.update({'partlist': reserved_list})

        print("Creating maintenance record...")
        maintenance_result = requests.request('POST', maintenanceAPI, json = data).json()
        maintenance_code = maintenance_result['code']

        amqp_setup.check_setup()

        # Successfully scheduled maintenance
        if maintenance_code == 201:

            # 3. Order Parts
            if len(missing_list):
                print("Ordering missing parts...")
                orderParts(missing_list)

            # 4. Send out Notification
            print("Sending Notification...")
            amqp_setup.channel.basic_publish(
                exchange=amqp_setup.exchangename, 
                routing_key="schedule.maintenance", 
                body=json.dumps(data), 
                properties=pika.BasicProperties(delivery_mode = 2)
            )

            print("Completed operations")
            return jsonify({
                "code": maintenance_code,
                "message": f"Maintenance for {eqp_id} has been scheduled on {schedule_date}"
            }), maintenance_code
        
        # Maintenance record already exist
        elif maintenance_code == 400:
            
            # Return reserved parts (RabbitMQ)
            print("Returning reserved parts...")
            returnParts(reserved_list)

            return jsonify({
                "code": maintenance_code,
                "message": "Maintenance record already exists"
            })
        
        else:
            return jsonify({
                "code": maintenance_code,
                "message": maintenance_result['message']
            })




    
@app.route("/start_maintenance/<string:maintenance_id>", methods=['PUT'])
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

    eqp_result = requests.request('PUT', f'{equipmentAPI}/{eqp_id}', json = eqp_data).json()

    maintenance_result = requests.request('PUT', f'{maintenanceAPI}/{maintenance_id}', json = maintenance_data).json()
    maintenance_code = maintenance_result['code']

    if maintenance_code == 201:
        executeMaintenance(data, start=True)

        return jsonify({
            "code": maintenance_code,
            "message": "Maintenance successfully started"
        }),maintenance_code
    
    else:
        return jsonify({
            "code": maintenance_code,
            "message": maintenance_result['message']
        })




@app.route("/request_parts/<string:maintenance_id>", methods=['PUT'])
def requestParts(maintenance_id):
    '''
    1. Reserve Parts (HTTP)
    2. Send update maintenance (Partlist)
    '''
    data = request.get_json()
    requested_parts = data["req_partlist"]
    current_parts = data["partlist"]

    print(current_parts)

    result = reserveParts(requested_parts)
    
    if type(result) != str:
        reserved_list, missing_list = result
        print(reserved_list)
        print(missing_list)

        additional_parts_dict = {}

        # Parts is available and reserved
        if len(reserved_list):
            for part in reserved_list:
                additional_parts_dict[part['_id']] = int(part['Qty'])

            for part in current_parts:
                if part['_id'] in additional_parts_dict:
                    part['Qty'] = str(int(part['Qty']) + additional_parts_dict[part['_id']])
                    del additional_parts_dict[part['_id']]

        # New additional parts
        if len(additional_parts_dict.keys()):
            for part in reserved_list:
                if part['_id'] in additional_parts_dict.keys():
                    current_parts.append(part)
        

        update = {
            'partlist': current_parts
        }

        print(current_parts)
        
        maintenance_result = requests.request('PUT', f'{maintenanceAPI}/{maintenance_id}', json = update).json()
        maintenance_code = maintenance_result['code']

        amqp_setup.check_setup()

        # Successfully scheduled maintenance
        if maintenance_code == 201:

            # 3. Order Parts
            if len(missing_list):
                orderParts(missing_list)

            return jsonify({
                "code": maintenance_code,
                "data": current_parts
            }), maintenance_code
        
        # Maintenance record failure
        else:
            
            # Return reserved parts (RabbitMQ)
            returnParts(reserved_list)

            return jsonify({
                "code": maintenance_code,
                "message": maintenance_result['message']
            })



@app.route("/end_maintenance/<string:maintenance_id>", methods=['PUT'])
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
            part['Qty'] = int(part['Qty']) - int(return_parts_dict[part['_id']])
            part['Qty'] = str(part['Qty'])

    eqp_data = {
        "equipment_status": eqp_status,
        "last_maintained": end_datetime
    }

    maintenance_data = {
        "end_datetime": end_datetime,
        "status": maintenance_status,
        "description": description,
        "partlist": part_list
    }

    maintenance_result = requests.request('PUT', f'{maintenanceAPI}/{maintenance_id}', json = maintenance_data).json()
    maintenance_code = maintenance_result['code']

    eqp_result = requests.request('PUT', f'{equipmentAPI}/{eqp_id}', json = eqp_data).json()

    if maintenance_code == 201:
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
        })




def reserveParts(partList):
    data = {
        "partList": partList
    }

    parts_result = requests.request('PUT', f'{inventoryAPI}/reserve', json=data).json()
    print("Part Results: ", parts_result)
    code = parts_result["code"] 

    # Parts sucessfully reserved
    if code == 200:
        reservedList = parts_result['data']['res_part_list']
        missingList = parts_result['data']['procurement_part_list']

        return (reservedList, missingList)
    
    else:
        return code['message']




def returnParts(partList):

    amqp_setup.check_setup()

    amqp_setup.channel.basic_publish(
        exchange=amqp_setup.exchangename, 
        routing_key="return.parts", 
        body=json.dumps(partList), 
        properties=pika.BasicProperties(delivery_mode = 2)
    )


def orderParts(partList):

    amqp_setup.check_setup()

    amqp_setup.channel.basic_publish(
        exchange=amqp_setup.exchangename, 
        routing_key="order.parts", 
        body=json.dumps(partList), 
        properties=pika.BasicProperties(delivery_mode = 2)
    )


def executeMaintenance(maintenanceData, start):
    if start:
        routing_key = "maintenance.start"
    else:
        routing_key = "maintenance.end"

    maintenanceData['start'] = start

    amqp_setup.check_setup()

    amqp_setup.channel.basic_publish(
        exchange=amqp_setup.exchangename, 
        routing_key=routing_key, 
        body=json.dumps(maintenanceData), 
        properties=pika.BasicProperties(delivery_mode = 2)
    )


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)