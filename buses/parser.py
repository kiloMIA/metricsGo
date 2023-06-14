import requests


def parse(route_number):
    link = 'http://45.135.131.226/api/buscoordinates/{route}'.format(route = route_number)
    r = requests.get(link)
    res_json = r.json()
    print(res_json)
    return res_json


