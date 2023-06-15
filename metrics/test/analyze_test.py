import unittest

import unittest
from unittest import TestCase
from ..analyzer import calculate


class TestCalculate(TestCase):
    def test_calculate_temperature(self):
        dicti = {
            'sensor1': {
                'temperature': 25,
                'humidity': 50
            },
            'sensor2': {
                'temperature': 30,
                'humidity': 60
            },
            'sensor3': {
                'temperature': 35,
                'humidity': 70
            }
        }

        result = calculate(dicti, 'temperature')

        expected_result = (30.0, 60.0)
        self.assertEqual(result, expected_result)

    def test_calculate_pol(self):
        dicti = {
            'sensor1': {
                'co2': 400,
                'pm25': 10
            },
            'sensor2': {
                'co2': 500,
                'pm25': 15
            },
            'sensor3': {
                'co2': 600,
                'pm25': 20
            }
        }

        result = calculate(dicti, 'pollution')

        expected_result = (500.0, 15.0)
        self.assertEqual(result, expected_result)


if __name__ == '__main__':
    unittest.main()
