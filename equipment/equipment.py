from flask import Flask, request, jsonify
from flask_cors import CORS
from get_database import get_database
from datetime import datetime
from bson import ObjectId

app = Flask(__name__)

DATABASE = get_database()
EQUIPMENT_COLLECTION = DATABASE["equipment"]
CORS(app)

def find_equipment(equipment_id):
    equipment_obj = EQUIPMENT_COLLECTION.find_one({"_id": equipment_id})
    equipment_obj = json(equipment_obj)

    return equipment_obj

def json(doc):
    if doc:
        doc["_id"] = str(doc["_id"])
        if type(doc["last_maintained"]) == datetime:
            doc["last_maintained"] = datetime.strftime(doc["last_maintained"], "%d-%m-%Y %H:%M:%S")
    
    return doc 


# GET ALL EQUIPMENT
@app.route("/equipment")
def get_all():
    
    all_equipment = EQUIPMENT_COLLECTION.find()
    equipmentList = []
    
    for eqmt in all_equipment:
        equipmentList.append(json(eqmt))
    
    if len(equipmentList):
        return jsonify({
            "code": 200,
            "data": {
                "equipment": [equipment for equipment in equipmentList]
            }
        }), 200
    
    return jsonify(
        {
            "code": 404,
            "message": "No equipment record"
        }
    ), 404

# GET SINGLE EQUIPMENT 
@app.route("/equipment/<string:equipment_id>")
def query_equipment(equipment_id):
    equipment_id = ObjectId(equipment_id)
    equipment_obj = find_equipment(equipment_id)

    # equipment record found
    if equipment_obj:
        return jsonify({
            "code": 200,
            "data": equipment_obj
        }), 200
    
    return jsonify(
        {
            "code": 404,
            "message": "No equipment record found."
        }
    ), 404


# CREATE EQUIPMENT 
@app.route("/equipment", methods=["POST"])
def create_equipment():
    data = request.get_json()
    
    try:
        EQUIPMENT_COLLECTION.insert_one(data)
    except: 
        return jsonify(
            {
                "code": 500,
                "data": {
                    "data": data
                },
                "message": "An error occurred creating the equipment."

            }
        ), 500
    
    # success
    return jsonify({
                "code": 200,
                "data": {
                    "message": "Successfully created"
                }
            }), 200


# UPDATE EQUIPMENT 
@app.route("/equipment/<string:equipment_id>", methods=["PUT"])
def update_equipment(equipment_id):
    equipment_id = ObjectId(equipment_id)
    data = request.get_json()
    data["last_maintained"] = datetime.strptime(data["last_maintained"], '%d-%m-%Y %H:%M:%S')
    
    try:
        EQUIPMENT_COLLECTION.update_one({"_id": equipment_id}, {"$set": data})
    except: 
        return jsonify(
            {
                "code": 500,
                "data": {
                    "equipment_name": equipment_id
                },
                "message": "An error occurred updating the equipment record."
            }
        ), 500

    equipment_obj = find_equipment(equipment_id)
    return jsonify({
                "code": 201,
                "data": {
                    "message": equipment_obj
                }
            }), 201

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=4999, debug=True)