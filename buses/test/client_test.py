import unittest
from unittest.mock import patch
from ..client import BusesServiceServicer
from ..buses_pb2 import BusRequest
from grpc import RpcContext


class TestBusesServiceServicer(unittest.TestCase):

    @patch('buses.client')
    def test_RequestBus(self, mock_parse):
        mock_parse.return_value = {'bus1': {'longitude': 123.45, 'latitude': 67.89}}

        servicer = BusesServiceServicer()

        context = RpcContext(None, None, None)

        request = BusRequest(BusNumber=37)

        response = servicer.RequestBus(request, context)

        mock_parse.assert_called_once_with('37')

        self.assertEqual(response.longitude, 123.45)
        self.assertEqual(response.latitude, 67.89)


if __name__ == '__main__':
    unittest.main()
