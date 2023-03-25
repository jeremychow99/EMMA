import logging
import grpc
import sys
import requests
import json

import controller
sys.path.append("./proto")
from proto import equipment_pb2
from proto import equipment_pb2_grpc

from concurrent import futures

def autoSchedule():
     # get list of technicians

     return

class Scheduler(equipment_pb2_grpc.SchedulerServicer):
        def ScheduleMaintenance(self, request, context):
            print(type(request.name))
            # TO-DO: Place function here to schedule maintenance
            # another function to clear any planned maintenance for the same equipment after
            req = requests.get('http://127.0.0.1:5000/maintenance').json()
            scheduled_eqp_list = []
        
            data = req['data']['maintenance']  
            for entry in data:
                 eqp_id = entry['equipment_id'] 
                 print(eqp_id)
                 scheduled_eqp_list.append(eqp_id)

            if request.name not in scheduled_eqp_list:
                print("id: " + request.name + " is not scheduled for maintenance")
            return equipment_pb2.Response(status="ok")
            

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    equipment_pb2_grpc.add_SchedulerServicer_to_server(Scheduler(), server)
    server.add_insecure_port('[::]:50051')
    
    server.start()
    print("gRPC starting on port 50051")
    server.wait_for_termination()
if __name__ == '__main__':
    serve()