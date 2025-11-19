import os
import pika
import time
from loguru import logger

class RabbitMQPublisher:
    def __init__(self):
        rabbitmq_url = os.getenv("RABBITMQ_URL")
        params = pika.URLParameters(rabbitmq_url)

        retries = 0
        max_retries = 30

        while retries < max_retries:
            try:
                logger.info(f"Connecting to RabbitMQ: {rabbitmq_url}")
                self.connection = pika.BlockingConnection(params)
                self.channel = self.connection.channel()
                logger.success("Connected to RabbitMQ successfully")
                break
            except Exception as e:
                retries += 1
                logger.warning(f"RabbitMQ not ready ({retries}/{max_retries}) - retrying...")
                time.sleep(2)

        if retries >= max_retries:
            raise Exception("Failed to connect to RabbitMQ after retries")

        self.channel.queue_declare(queue=os.getenv("QUEUE_NAME"), durable=True)
