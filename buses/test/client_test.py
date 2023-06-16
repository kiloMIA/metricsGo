import unittest
from unittest.mock import patch, MagicMock
from ..client import BusesServiceServicer
from ..buses_pb2 import BusRequest, BusResponse


class TestBusesServiceServicer(unittest.TestCase):
    @patch('buses.parse')
    @patch('buses.send_bus_response')
    def test_RequestBus(self, mock_send_bus_response, mock_parse):
        mock_parse.return_value = {
            'bus1': {'longitude': 71.49366, 'latitude': 51.15155},
            'bus2': {'longitude': 71.50432, 'latitude': 51.15543},
            'bus3': {'longitude': None, 'latitude': None},
        }
        servicer = BusesServiceServicer()
        request = MagicMock()
        request.BusNumber = 123

        response = servicer.RequestBus(request, MagicMock())

        mock_parse.assert_called_once_with(123)
        mock_send_bus_response.assert_called_with(BusResponse(longitude=71.49366, latitude=51.15155))
        mock_send_bus_response.assert_called_with(BusResponse(longitude=71.50432, latitude=51.15543))
        self.assertIsInstance(response, BusResponse)


if __name__ == '__main__':
    unittest.main()
