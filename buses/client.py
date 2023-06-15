import logging
from concurrent import futures
import grpc

from buses.producer import send_bus_response
from parser import parse
from buses_pb2_grpc import add_BusServiceServicer_to_server, BusServiceServicer
from buses_pb2 import BusResponse


class BusesServiceServicer(BusServiceServicer):
    def RequestBus(self, request, context):
        route_number = request.BusNumber
        location = parse(route_number)
        logging.info(location)
        for bus in location:
            send_bus_response(BusResponse(longitude=location[bus]['longitude'], latitude=location[bus]['latitude']))


def serve():
    logging.basicConfig(level=logging.INFO)
    logging.info('Starting gRPC server...')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_BusServiceServicer_to_server(BusesServiceServicer(), server)
    server.add_insecure_port('[::]:50054')
    server.start()
    logging.info('server has started...')
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
