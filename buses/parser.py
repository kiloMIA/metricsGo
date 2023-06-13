import aiohttp
import asyncio


async def parse(route_number):
    async with aiohttp.ClientSession() as session:
        link = 'http://45.135.131.226/api/buscoordinates/{route_number}'.format(route_number=route_number)
        async with session.get(link) as resp:
            res_json = await resp.json()
        return res_json


async def main():
    await parse(37)


loop = asyncio.get_event_loop()
loop.run_until_complete(main())
