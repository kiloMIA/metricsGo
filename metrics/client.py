from concurrent import futures

import grpc
from analyzer import analyze

from metrics_pb2 import TemperatureResponse
from metrics_pb2_grpc import add_MetricsServiceServicer_to_server, MetricsServiceServicer


# logging.basicConfig(format='%(asctime)s:%(name)s : %(message)s', level=logging.DEBUG)
# log = logging.getLogger(__name__)
#
# socket = 'localhost:50051'
# channel = grpc.insecure_channel(socket)
# client = mgrpc.MetricsServiceStub(channel)
# log.info(f"Connected to the Server : {socket}")


class TemperatureServiceServicer(MetricsServiceServicer):
    async def CalculateTemperature(self, request, context):
        city = request.city
        req_type = request.type
        district_list = await analyze(city, req_type)
        for d in district_list:
            district = d['district']
            temperature = d['temperature']
            humidity = d['humidity']
            return TemperatureResponse(city=city, district=district, temperature=temperature, humidity=humidity)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_MetricsServiceServicer_to_server(TemperatureServiceServicer(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
