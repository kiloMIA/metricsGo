import requests


def parse(route_number):
    link = 'http://opendata.kz/api/sensor/getListWithLastHistory?cityId={city_id}'.format(city_id=city_id)
    r = requests.get(link)
    res_json = r.json()
    return res_json


