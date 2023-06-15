import pika
import json
from buses_pb2 import BusResponse


def send_bus_response(bus_response):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()
    channel.queue_declare(queue='bus_queue')

    message = json.dumps({"longitude": bus_response.longitude, "latitude": bus_response.latitude})

    channel.basic_publish(exchange='', routing_key='bus_queue', body=message.encode())
    print("Sent bus response")

    connection.close()
