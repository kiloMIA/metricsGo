import aiohttp


async def parse(city_id: int, req_type: str):
    async with aiohttp.ClientSession() as session:
        link = 'http://opendata.kz/api/sensor/getListWithLastHistory?cityId={city_id}'.format(city_id=city_id)
        async with session.get(link) as resp:
            res_json = await resp.json()

    metrics_dict = {}
    for sensor in res_json['sensors']:
        if req_type == "temperature":
            metrics_dict.update({
                sensor['name']: {
                    'temperature': sensor['history'][0]['data']['field3'],
                    'humidity': sensor['history'][0]['data']['field5']
                }})
        else:
            metrics_dict.update({
                sensor['name']: {
                    'co2': sensor['history'][0]['data']['field1'],
                    'pm25': sensor['history'][0]['data']['field2']
                }})

    return metrics_dict
