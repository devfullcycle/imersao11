from confluent_kafka import Producer
import socket
import logging

logger = logging.getLogger(__name__)

conf = {
    'bootstrap.servers': 'host.docker.internal:9094',
    'client.id': socket.gethostname(),
}


def acked(err, msg):
    if err is not None:
        print(f"Failed to deliver message: {str(msg)}: {str(err)}")
    else:
        print(f"Message produced: {str(msg)}")


def publish_message(topic, value, key='key'):
    producer = Producer(conf)
    producer.produce(topic, value=value, key=key, callback=acked)
    producer.poll(3)


def safe_publish_message(topic, value, key='key'):
    try:
        publish_message(topic, value, key)
    except Exception as e:
        logger.exception(e)
