import pika
from os import environ

# hostname = environ.get('rabbit_host') or 'localhost'
# port = environ.get('rabbit_port') or 5672
hostname = 'rabbitmq'
port = 5672


# connect to the broker and set up a communication channel in the connection
connection = pika.BlockingConnection(
    pika.ConnectionParameters(
        host=hostname, port=port,
        heartbeat=3600, blocked_connection_timeout=3600, 
))
    

channel = connection.channel()

exchangename= "emma_topic" 
exchangetype= "topic" 
channel.exchange_declare(exchange=exchangename, exchange_type=exchangetype, durable=True)

############   Maintenance queue    #############
queue_name = 'Maintenance_SMS' 
channel.queue_declare(queue=queue_name, durable=True)

routing_key = 'schedule.maintenance' 

channel.queue_bind(exchange=exchangename, queue=queue_name, routing_key=routing_key) 

queue_name = 'Maintenance_Email' 
channel.queue_declare(queue=queue_name, durable=True)

routing_key = 'schedule.maintenance' 

channel.queue_bind(exchange=exchangename, queue=queue_name, routing_key=routing_key) 


############   Execute Maintenance queue    #############
queue_name = 'Execute_Maintenance' 
channel.queue_declare(queue=queue_name, durable=True)

routing_key = 'maintenance.*' 

channel.queue_bind(exchange=exchangename, queue=queue_name, routing_key=routing_key) 

"""
This function in this module sets up a connection and a channel to a local AMQP broker,
and declares a 'topic' exchange to be used by the microservices in the solution.
"""
def check_setup():
    # The shared connection and channel created when the module is imported may be expired, 
    # timed out, disconnected by the broker or a client;
    # - re-establish the connection/channel is they have been closed
    global connection, channel, hostname, port, exchangename, exchangetype

    if not is_connection_open(connection):
        connection = pika.BlockingConnection(pika.ConnectionParameters(host=hostname, port=port, heartbeat=3600, blocked_connection_timeout=3600))
    if channel.is_closed:
        channel = connection.channel()
        channel.exchange_declare(exchange=exchangename, exchange_type=exchangetype, durable=True) ###


def is_connection_open(connection):
    # For a BlockingConnection in AMQP clients,
    # when an exception happens when an action is performed,
    # it likely indicates a broken connection.
    # So, the code below actively calls a method in the 'connection' to check if an exception happens
    try:
        connection.process_data_events()
        return True
    except pika.exceptions.AMQPError as e:
        print("AMQP Error:", e)
        print("...creating a new connection.")
        return False
