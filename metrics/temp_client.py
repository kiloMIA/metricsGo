from concurrent import futures
import logging
import asyncio
import json
import grpc
from analyzer import analyze

from metrics_pb2 import TemperatureResponse, PollutionResponse
from metrics_pb2_grpc import add_MetricsServiceServicer_to_server, MetricsServiceServicer


# logging.basicConfig(format='%(asctime)s:%(name)s : %(message)s', level=logging.DEBUG)
# log = logging.getLogger(__name__)
#
# socket = 'localhost:50051'
# channel = grpc.insecure_channel(socket)
# client = mgrpc.MetricsServiceStub(channel)
# log.info(f"Connected to the Server : {socket}")


class TemperatureServiceServicer(MetricsServiceServicer):
    def RequestTemp(self, request, context):
        city = request.city
        reqType = request.type
        district_list = analyze(city, reqType)

        logging.info(district_list)
        for district in district_list:
            return TemperatureResponse(
                city=city,
                district=district['district'],
                temperature=district['temperature'],
                humidity=district['humidity']
            )
    def RequestPol(self, request, context):
        city = request.city
        reqType = request.type
        district_list = analyze(city, reqType)
        logging.info(district_list)
        for district in district_list:
            return PollutionResponse(
                city=city,
                co2=district['co2'],
                district=district['district'],
                pm25=district['pm25']
            )


def serve():
    logging.basicConfig(level=logging.INFO)
    logging.info('Starting gRPC server...')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_MetricsServiceServicer_to_server(TemperatureServiceServicer(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    logging.info('server has started...')
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
