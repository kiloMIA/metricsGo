import unittest
from unittest.mock import patch, MagicMock
from ..metrics_pb2 import  TemperatureResponse, PollutionResponse
from ..metrics_pb2_grpc import MetricsServiceServicer

class TestTemperatureServiceServicer(unittest.TestCase):
    @patch('metrics.analyzer.analyze')
    @patch('metrics.producer.send_to_queue')
    def test_RequestTemp(self, mock_send_to_queue, mock_analyze):
        mock_analyze.return_value = [
            {
                'district': 'district1',
                'temperature': 25.0,
                'humidity': 70.0
            },
            {
                'district': 'district2',
                'temperature': 28.5,
                'humidity': 65.0
            }
        ]
        servicer = MetricsServiceServicer()
        request = MagicMock()
        request.city = 'city1'
        request.type = 'type1'

        response = servicer.RequestTemp(request, MagicMock())

        mock_analyze.assert_called_once_with('city1', 'type1')
        mock_send_to_queue.assert_called_with(
            {
                'city': 'city1',
                'district': 'district1',
                'temperature': 25.0,
                'humidity': 70.0
            },
            'temperature_queue'
        )
        mock_send_to_queue.assert_called_with(
            {
                'city': 'city1',
                'district': 'district2',
                'temperature': 28.5,
                'humidity': 65.0
            },
            'temperature_queue'
        )
        self.assertIsInstance(response, TemperatureResponse)

    @patch('metrics.analyzer.analyze')
    @patch('metrics.producer.send_to_queue')
    def test_RequestPol(self, mock_send_to_queue, mock_analyze):
        mock_analyze.return_value = [
            {
                'district': 'district1',
                'co2': 100.0,
                'pm25': 50.0
            },
            {
                'district': 'district2',
                'co2': 120.0,
                'pm25': 60.0
            }
        ]
        servicer = MetricsServiceServicer()
        request = MagicMock()
        request.city = 'city1'
        request.type = 'type1'

        response = servicer.RequestPol(request, MagicMock())

        mock_analyze.assert_called_once_with('city1', 'type1')
        mock_send_to_queue.assert_called_with(
            {
                'city': 'city1',
                'district': 'district1',
                'co2': 100.0,
                'pm25': 50.0
            },
            'pollution_queue'
        )
        mock_send_to_queue.assert_called_with(
            {
                'city': 'city1',
                'district': 'district2',
                'co2': 120.0,
                'pm25': 60.0
            },
            'pollution_queue'
        )
        self.assertIsInstance(response, PollutionResponse)

if __name__ == '__main__':
    unittest.main()
