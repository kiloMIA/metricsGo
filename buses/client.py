from concurrent import futures
import grpc
from parser import parse
from buses_pb2_grpc import add_BusServiceServicer_to_server, BusServiceServicer
from buses_pb2 import BusResponse


class BusesServiceServicer(BusServiceServicer):
    def FindBus(self, request, context):
        route_number = request.BusNumber
        location = parse(route_number)
        for bus in location:
            return BusResponse(longitude=location[bus]['longitude'], latitude=location[bus]['latitude'])


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_BusServiceServicer_to_server(BusesServiceServicer(), server)
    server.add_insecure_port('[::]:50054')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
