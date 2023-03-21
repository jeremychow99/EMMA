import requests
import json

from pymongo import MongoClient
from flask import Flask, request, jsonify
from flask_cors import CORS
from dotenv import load_dotenv
from bson import ObjectId
from invokes import invoke_http
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
# GET ALL PARTS
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

#REQUEST FOR A PART AND UPDATE DB IF AVAILABLE, IF NOT INVOKE EMAIL MS
@app.route("/inventory/<string:inventory_id>&<int:quantity>/reserve", methods = ['PUT']) 
def request_part(inventory_id,quantity):
    # Get part_id and quantity from request
    req_part_id = ObjectId(inventory_id)
    req_quantity = quantity
    
    # Check if required quantity for part is available from inventory database
    part = collection.find_one({'_id': req_part_id})
    if part['Qty'] >= req_quantity:
        #If yes, reserve quantity by updating the inventory database
        new_qty = part['Qty'] - req_quantity
        collection.update_one({'_id': req_part_id}, {'$set': {'Qty': new_qty}}) 
        return jsonify({
                "code": 200,
                "message": "Quantity reserved"
            }), 200
        
    else:
        # If no, determine missing quantity and trigger procurement process (triggers notification microservice)
        available_quantity = part['Qty']
        missing_quantity = req_quantity - available_quantity
        
        invoke_http(f"/email/{missing_quantity}", method='PUT')
        return jsonify({
            "code": 400,
            "message" : f'Not available. Missing Quantity: {missing_quantity}. Procurement initiated.'
        }), 400

#RETURN PARTS TO DB
@app.route("/inventory/<string:inventory_id>&<int:quantity>/return", methods = ['PUT'])
def return_parts(inventory_id, quantity):
    # Get part_id and quantity from request
    req_part_id = ObjectId(inventory_id)
    req_quantity = quantity

    # Add quantity back to inventory database
    part = collection.find_one({'_id': req_part_id})
    new_qty = part['Qty'] + req_quantity
    collection.update_one({'_id': req_part_id}, {'$set': {'Qty': new_qty}})

    return jsonify({
        "code": 200,
        "message": "Parts added back to inventory."
    }), 200
 
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=True)





