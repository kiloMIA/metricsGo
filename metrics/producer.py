import json

import pika


def send_to_queue(data, queue_name):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='rabbitmq'))
    channel = connection.channel()
    channel.queue_declare(queue=queue_name)

    # Convert the response to a JSON string
    message = json.dumps(data)

    channel.basic_publish(exchange='', routing_key=queue_name, body=message)
    print(f"Sent message to {queue_name}")

    connection.close()