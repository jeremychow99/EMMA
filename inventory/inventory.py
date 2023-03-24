from pymongo import MongoClient
from flask import Flask, request, jsonify
from flask_cors import CORS
from dotenv import load_dotenv
from bson import ObjectId
import os


load_dotenv()

############# Configurations ##############
app = Flask(__name__)


DB_URI = os.getenv('DB_URI')
DB_NAME = 'Inventory'
COLLECTION_NAME = 'inventory'

client = MongoClient(DB_URI)
db = client[DB_NAME]
collection = db[COLLECTION_NAME]

CORS(app)

##########################################
#GET ALL PARTS
@app.route("/inventory")
def get_all():
    
    all_parts = collection.find()
    partList = []
    
    for part in all_parts:
        part['_id'] = str(part['_id'])
        partList.append(part)
    
    if len(partList):
        return jsonify({
            "code": 200,
            "data": {
                "parts": partList
            }
        }), 200
    
    return jsonify(
        {
            "code": 404,
            "message": "No Parts Found."
        }
    ), 404

#REQUEST FOR PARTS AND UPDATE DB IF AVAILABLE
@app.route("/inventory/reserve", methods = ['PUT']) 
def reserve_parts():
    # Get parts and quantities from JSON request body
    data = request.get_json()
    parts = data['partList']
    
    #reserved parts list and procurement parts list to return
    res_part_list = []
    procurement_part_list = []
    error_part_list = []

    for each in parts:
        # Get part_id and quantity from request
        req_part_id = ObjectId(each.get('_id'))
        req_quantity = int(each.get('Qty'))
        req_partname = str(each.get('PartName'))

        try:
            # Check if required quantity for part is available from inventory database
            part = collection.find_one({'_id': req_part_id})
            if part['Qty'] >= req_quantity:
                #If yes, reserve quantity by updating the inventory database
                new_qty = part['Qty'] - req_quantity
                collection.update_one({'_id': req_part_id}, {'$set': {'Qty': new_qty}}) 
                res_part_list.append({
                    "PartName" : req_partname,
                    "ReservedQty": req_quantity,
                    "_id": str(req_part_id)
                    })

            else:
                # If no, determine missing quantity and trigger procurement process (triggers notification microservice)
                available_quantity = part['Qty']
                missing_quantity = req_quantity - available_quantity
                collection.update_one({'_id': req_part_id}, {'$set': {'Qty': 0}})
                res_part_list.append({
                    "PartName" : req_partname,
                    "Qty": req_quantity,
                    "_id": str(req_part_id)
                    })
                procurement_part_list.append({
                    "PartName" : req_partname,
                    "Qty": missing_quantity,
                    "_id": str(req_part_id)})

        except:
            error_part_list.append(f'{req_part_id}')
            continue

    if len(error_part_list):
        return jsonify({
                "code": 500,
                "message": "Error occurred when updating inventory.",
                "data": {
                    "error_part_list": error_part_list,
                    "res_part_list": res_part_list,
                    "procurement_part_list": procurement_part_list
                }
            }), 500

    return jsonify({
                "code": 200,
                "message": "Part(s) reserved.",
                "data": {
                    "res_part_list": res_part_list,
                    "procurement_part_list": procurement_part_list
                }
            }), 200

#RETURN PARTS TO DB
@app.route("/inventory/return", methods = ['PUT'])
def return_parts():
    # Get parts and quantities from JSON request body
    data = request.get_json()
    parts = data['partList']

    #returned parts list to return
    returned_part_list = []
    error_part_list = []

    for each in parts:
        # Get part_id and quantity from request
        req_part_id = ObjectId(each.get('_id'))
        req_quantity = int(each.get('Qty'))
        req_partname = str(each.get('PartName'))


        try:
            # Add quantity back to inventory database
            part = collection.find_one({'_id': req_part_id})
            new_qty = part['Qty'] + req_quantity
            collection.update_one({'_id': req_part_id}, {'$set': {'Qty': new_qty}})
            returned_part_list.append({
                "PartName" : req_partname,
                "ReturnedQty": req_quantity,
                "_id": str(req_part_id)
                })

        except:
            error_part_list.append(f'{req_part_id}')
            continue

    if len(error_part_list):
        return jsonify({
                "code": 500,
                "message": "Error: Part(s) do not exist.",
                "data": {
                    "error_part_list": error_part_list,
                    "returned_part_list": returned_part_list
                }
            }), 500


    return jsonify({
        "code": 200,
        "message": "Parts added back to inventory.",
        "data": {
                    "returned_part_list": returned_part_list
        }
    }), 200
 
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=True)





