from pymongo import MongoClient
from flask import Flask, request, jsonify
from flask_cors import CORS
from datetime import datetime
from bson import ObjectId
import os
from dotenv import load_dotenv

load_dotenv()

############# Configurations ##############
app = Flask(__name__)


DB_URI = os.getenv('DB_URI')
DB_NAME = 'Maintenance'
COLLECTION_NAME = 'maintenance'

client = MongoClient(DB_URI)
db = client[DB_NAME]
collection = db[COLLECTION_NAME]

CORS(app)


############# Helper Function ##############
def json(doc):
    if doc:
        # Convert ID to string
        doc['_id'] = str(doc['_id'])

        # Convert Datetime to string (dd/mm/yyyy hh:mm:ss)
        datetime_keys = ['schedule_datetime', 'start_datetime', 'end_datetime']
        for key in datetime_keys:
            if key in doc:
                doc[key] = datetime.strftime(doc[key], '%d-%m-%Y %H:%M:%S')

    return doc        

# type(maintenance_id) == ObjectId
def find_maintenance(maintenance_id):

    maintenance_obj = collection.find_one({"_id": maintenance_id})
    maintenance_obj = json(maintenance_obj)

    return maintenance_obj



############# API Routes ##############

# Retrieve all maintenance records
@app.route("/maintenance")
def query_all_maintenance():

    maintenance_list = []

    # Iterate through all documents in collection via Cursor instance
    for maintenance in collection.find():
        maintenance = json(maintenance)
        maintenance_list.append(maintenance)

    # If any document exist
    if len(maintenance_list):
        return jsonify({
            "code": 200,
            "data": {
                "maintenance": [maintenance for maintenance in maintenance_list]
            }
        }), 200
    
    return jsonify(
        {
            "code": 404,
            "message": "No maintenance record."
        }
    ), 404



# Retrieve single maintenance record by id
@app.route("/maintenance/<string:maintenance_id>")
def query_maintenance(maintenance_id):

    maintenance_id = ObjectId(maintenance_id)

    maintenance_obj = find_maintenance(maintenance_id)

    # Maintenance Record found
    if maintenance_obj:
        return jsonify({
            "code": 200,
            "data": maintenance_obj
        }), 200
    
    return jsonify(
        {
            "code": 404,
            "message": "No maintenance record found."
        }
    ), 404



# Update single maintenance record
@app.route("/maintenance/<string:maintenance_id>", methods=['PUT'])
def update_maintenance(maintenance_id):

    maintenance_id = ObjectId(maintenance_id)

    data = request.get_json()

    try:
        # Update maintenance record
        result = collection.update_one({"_id":maintenance_id}, {"$set": data})
        print("Records Found: ", result.matched_count)
        print("Records Modified: ", result.modified_count)
        
        # No matching document
        if result.matched_count == 0:
            return jsonify({
                "code": 404,
                "data": {
                    "maintenance_id": str(maintenance_id)
                },
                "message": "No maintenance record found."
            }), 404
            
    except:
        return jsonify({
            "code": 500,
            "data": {
                "maintenance_id": str(maintenance_id)
            },
            "message": "Error occurred when updating maintenance record."
        }), 500 
    
    # Successfully updated maintenance record
    maintenance_obj = find_maintenance(maintenance_id)

    return jsonify({
        "code": 201,
        "data": maintenance_obj
    }), 201



# Create maintenance record
@app.route("/maintenance", methods=['POST'])
def schedule_maintenance():

    data = request.get_json()
    eqp_id = data["equipment_id"]

    if "schedule_datetime" not in data:
        data["schedule_datetime"] = datetime.now()

    else:
        data["schedule_datetime"] = datetime.strptime(data["schedule_datetime"], '%d-%m-%Y %H:%M:%S')

    sched_datetime = data["schedule_datetime"]

    # Validate if there is a document with same equipment_id and schedule_datetime
    # Need a better way of determining if there is schedule for the equipment today (Can just match date?)
    maintenance_obj = collection.find_one({"equipment_id": eqp_id, "schedule_datetime": sched_datetime})

    # Maintenance exists
    if maintenance_obj:
        return jsonify({
            "code": 400,
            "data": maintenance_obj,
            "message": "Maintenance record already exists."
        }), 400
    

    try:
        # Add new maintenance record
        new_maintenance_id = collection.insert_one(data).inserted_id

    except:
        return jsonify({
            "code": 500,
            "data": data,
            "message": "Error occurred when creating maintenance record"
        }), 500


    maintenance_obj = find_maintenance(new_maintenance_id)
    return jsonify({
        "code": 201,
        "data": maintenance_obj
    }), 201











if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000, debug=True)

