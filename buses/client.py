from concurrent import futures
import grpc
from parser import parse
from buses_pb2_grpc import add_BusServiceServicer_to_server, BusServiceServicer
from buses_pb2 import BusResponse
# logging.basicConfig(format='%(asctime)s:%(name)s : %(message)s', level=logging.DEBUG)
# log = logging.getLogger(__name__)
#
# socket = 'localhost:50051'
# channel = grpc.insecure_channel(socket)
# client = mgrpc.MetricsServiceStub(channel)
# log.info(f"Connected to the Server : {socket}")


class BusesServiceServicer(BusServiceServicer):
    async def FindBus(self, request, context):
        route_number = request.BusNumber
        location = await parse(route_number)
        for bus in location:
            return BusResponse(longitude=location[bus]['longitude'], latitude=location[bus]['latitude'])


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_BusServiceServicer_to_server(BusesServiceServicer(), server)
    server.add_insecure_port('[::]:50053')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
