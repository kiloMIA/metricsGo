import unittest
from unittest import TestCase
from unittest.mock import patch
from ..parser import parse


class TestParse(TestCase):

    @patch('metrics.parser.requests.get')
    def test_parse(self, mock_get):
        mock_get.return_value.json.return_value = {
            'sensors': [
                {
                    'name': 'sensor1',
                    'history': [
                        {
                            'data': {
                                'field1': 'co2_value',
                                'field2': 'pm25_value',
                                'field3': 'temperature_value',
                                'field5': 'humidity_value'
                            }
                        }
                    ]
                }
            ]
        }

        result = parse(1, 'temperature')

        mock_get.assert_called_once_with('http://opendata.kz/api/sensor/getListWithLastHistory?cityId=1')

        expected_result = {
            'sensor1': {
                'temperature': 'temperature_value',
                'humidity': 'humidity_value'
            }
        }

        self.assertEqual(result, expected_result)


if __name__ == '__main__':
    unittest.main()
