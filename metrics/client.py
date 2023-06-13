from concurrent import futures
import logging
import asyncio
import json
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
    def RequestTemp(self, request, context):
        city = request.city
        #req_type = request.type
        req_type="temperature"
        district_list = asyncio.run(analyze(city, req_type))

        logging.info(district_list)
        for district in district_list:
            return TemperatureResponse(
                city=city,
                district=district['district'],
                temperature=district['temperature'],
                humidity=district['humidity']
            )
      

def serve():
    logging.basicConfig(level=logging.INFO)
    logging.info('Starting gRPC server...')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_MetricsServiceServicer_to_server(TemperatureServiceServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    logging.info('server has started...')
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
