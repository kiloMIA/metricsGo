import json

import pika


def send_temperature_response(temperature_response):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()
    channel.queue_declare(queue='temperature_queue')

    message = json.dumps({
        "city": temperature_response.city,
        "district": temperature_response.district,
        "temperature": temperature_response.temperature,
        "humidity": temperature_response.humidity
    })

    channel.basic_publish(exchange='', routing_key='temperature_queue', body=message.encode())
    print("Sent temperature response")

    connection.close()


def send_pollution_response(pollution_response):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()
    channel.queue_declare(queue='pollution_queue')

    message = json.dumps({
        "city": pollution_response.city,
        "district": pollution_response.district,
        "co2": pollution_response.co2,
        "pm25": pollution_response.pm25
    })

    channel.basic_publish(exchange='', routing_key='pollution_queue', body=message.encode())
    print("Sent pollution response")

    connection.close()
