import unittest
from unittest.mock import MagicMock, Mock
from ..producer import send_bus_response


class TestSendBusResponse(unittest.TestCase):
    def setUp(self):
        self.connection_mock = Mock()
        self.channel_mock = Mock()
        self.channel_mock.queue_declare.return_value = None
        self.connection_mock.channel.return_value = self.channel_mock

    def test_send_bus_response(self):
        bus_response = Mock()
        bus_response.longitude = 123.456
        bus_response.latitude = 78.901

        expected_message = '{"longitude": 123.456, "latitude": 78.901}'

        with unittest.mock.patch('pika.BlockingConnection', return_value=self.connection_mock):
            send_bus_response(bus_response)

        self.channel_mock.basic_publish.assert_called_once_with(exchange='', routing_key='bus_queue',
                                                                body=expected_message.encode())
        self.channel_mock.queue_declare.assert_called_once_with(queue='bus_queue')
        self.connection_mock.close.assert_called_once()


if __name__ == '__main__':
    unittest.main()
