import logging
import grpc
import sys

sys.path.append("./proto")
from proto import equipment_pb2
from proto import equipment_pb2_grpc

from concurrent import futures

class Scheduler(equipment_pb2_grpc.SchedulerServicer):
        def ScheduleMaintenance(self, request, context):
             print(str(request))
             # TO-DO: Place function here to schedule maintenance
             # another function to clear any planned maintenance for the same equipment after
             return equipment_pb2.Response(status="ok")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    equipment_pb2_grpc.add_SchedulerServicer_to_server(Scheduler(), server)
    server.add_insecure_port('[::]:50051')
    
    server.start()
    print("gRPC starting on port 50051")
    server.wait_for_termination()

serve()