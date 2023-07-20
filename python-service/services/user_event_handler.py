from typing import Any
import pika
import json

def emit_user_profile_update(user_id, new_data):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='rabbitmq-server'))
    channel = connection.channel()

    exchange_name = "user_updates"
    routing_key = "user.profile.update"

    channel.exchange_declare(exchange=exchange_name, exchange_type='topic', durable=True)

    new_data["id"] = user_id

    channel.basic_publish(
        exchange=exchange_name, 
        routing_key=routing_key,
        properties=pika.BasicProperties(delivery_mode=2)
    )
    print("%r sent to exchange %r with data: %r" % (routing_key, exchange_name, new_data))
    connection.close()








