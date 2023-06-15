import logging
from concurrent import futures
import grpc

from producer import send_bus_response
from parser import parse
from buses_pb2_grpc import add_BusServiceServicer_to_server, BusServiceServicer
from buses_pb2 import BusResponse


class BusesServiceServicer(BusServiceServicer):
    def RequestBus(self, request, context):
        route_number = request.BusNumber
        location = parse(route_number)

        for bus in location:
            try:
                bus_data = location[bus]
                if 'longitude' in bus_data and 'latitude' in bus_data and bus_data['longitude'] is not None and bus_data['latitude'] is not None:
                    bus_response = BusResponse(longitude=bus_data['longitude'], latitude=bus_data['latitude'])
                    logging.info(bus_response)
                    send_bus_response(bus_response)
                else:
                    logging.warning(f"Bus data for {bus} is missing longitude or latitude.")
            except Exception as e:
                logging.error(f"Error processing bus data for {bus}: {str(e)}")

        return BusResponse()  # Return an empty BusResponse



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
