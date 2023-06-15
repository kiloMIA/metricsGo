import unittest
from unittest import TestCase
from unittest.mock import patch
from buses.parser import parse


class TestParse(TestCase):

    @patch('buses.parser.requests.get')
    def test_parse(self, mock_get):
        mock_get.return_value.json.return_value = {'key': 'value'}

        result = parse('37')

        mock_get.assert_called_once_with('http://45.135.131.226/api/buscoordinates/37')

        self.assertEqual(result, {'key': 'value'})


if __name__ == '__main__':
    unittest.main()
