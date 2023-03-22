from flask import Flask, request, jsonify
import os
from dotenv import load_dotenv, environ
import requests

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
    eqpId = data["equipment_id"]
    scheduled_datetime = data["schedule_datetime"]
    requested_parts = data["partlist"]

    # 1. Reserve Parts (HTTP)
    result = reserveParts(requested_parts)

    if type(result) != str:

        reservedList, missingList = result

        # 2. Create Maintenance Record
        data['partList'] = reservedList

        maintenance_result = requests.request('POST', maintenanceAPI, json = data)
        maintenance_code = maintenance_result['code']

        # Successfully scheduled maintenance
        if maintenance_code == '201':

            # 3. Order Parts

            # 4. Send out Notification

            return jsonify({
                "code": maintenance_code,
                "message": f"Maintenance for {eqpId} has been scheduled on {scheduled_datetime}"
            })
        
        # Maintenance record already exist
        elif maintenance_code == '404':
            
            # Return reserved parts (RabbitMQ)

            return jsonify({
                "code": maintenance_code,
                "message": "Maintenance record already exists"
            })
        
        else:

            return jsonify({
                "code": maintenance_code,
                "message": maintenance_result['message']
            })




    
    




@app.route("/start_maintenance/<string:maintenanceId>")
def startMaintenance():
    '''
    1. Send update equipment (Status to maintenance)
    2. Send update maintenance (Start time, techID, Status)
    3. Send out Notification (Email + SMS)
    '''
    pass



@app.route("/request_parts/<string:maintenanceId>")
def requestParts():
    '''
    1. Reserve Parts (HTTP)
    2. Send update maintenance (Partlist)
    '''
    pass



@app.route("/end_maintenance/<string:maintenanceId>")
def endMaintenance():
    '''
    1. Send update maintenance (Last Service time, Status, Description)
    2. Send update equipment (Status to maintenance)
    3. Return Parts (RabbitMQ)
    3. Send out Notification (Email + SMS)
    '''
    pass




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
    pass


def orderParts(partList):
    pass
