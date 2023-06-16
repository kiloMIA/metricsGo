import unittest
from unittest.mock import Mock
from ..producer import send_to_queue

class TestSendToQueue(unittest.TestCase):
    def setUp(self):
        self.connection_mock = Mock()
        self.channel_mock = Mock()
        self.channel_mock.queue_declare.return_value = None
        self.connection_mock.channel.return_value = self.channel_mock

    def test_send_to_queue(self):
        data = {'key': 'value'}
        queue_name = 'test_queue'
        expected_message = '{"key": "value"}'

        with unittest.mock.patch('pika.BlockingConnection', return_value=self.connection_mock):
            send_to_queue(data, queue_name)

        self.channel_mock.basic_publish.assert_called_once_with(exchange='', routing_key=queue_name, body=expected_message)
        self.channel_mock.queue_declare.assert_called_once_with(queue=queue_name)
        self.connection_mock.close.assert_called_once()

if __name__ == '__main__':
    unittest.main()
