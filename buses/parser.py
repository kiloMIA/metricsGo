import aiohttp


async def parse(route_number):
    async with aiohttp.ClientSession() as session:
        link = 'http://45.135.131.226/api/buscoordinates/{route_number}'.format(route_number=route_number)
        async with session.get(link) as resp:
            res_json = await resp.json()
