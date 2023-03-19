import requests
import json
import os

from pymongo import MongoClient
from flask import Flask, request, jsonify
from flask_cors import CORS
from dotenv import load_dotenv

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

@app.route("/inventory")
def request_parts():
    # Get part_id and quantity from request
    part_id = request.args.get('Part_Id')
    quantity = int(request.args.get('Qty'))

    # Check if required quantity for part is available from inventory database
    part = collection.find_one({'Part_Id': part_id})
    if part['Qty'] >= quantity:
        # If yes, reserve quantity by updating the inventory database
        new_qty = part['Qty'] - quantity
        collection.update_one({'Part_id': part_id}, {'$set': {'Qty': new_qty}}) 
        return jsonify({
                "code": 200,
                "message": "Quantity reserved"
            }), 200
        
    else:
        # If no, determine missing quantity and trigger procurement process (triggers notification microservice)
        available_quantity = part['Qty']
        missing_quantity = quantity - available_quantity
        
        requests.post(f"/notifications")
        return jsonify({
            "code": 400,
            "message" : 'Not available. Missing Quantity. Procurement initiated.',
            "data": missing_quantity
        }), 400

@app.route("/inventory/<string:maintenance_id>")
def return_parts():
    # Get part_id and quantity from request
    part_id = request.args.get('Part_Id')
    quantity = int(request.args.get('Qty'))

    # Add quantity back to inventory database
    part = collection.find_one({'Part_Id': part_id})
    new_qty = part['Qty'] + quantity
    collection.update_one({'Part_Id': part_id}, {'$set': {'Qty': new_qty}})

    return jsonify({
        "code": 200,
        "message": "Parts added back to inventory."
    }), 200
 
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=True)


