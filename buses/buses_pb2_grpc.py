# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import buses_pb2 as buses__pb2


class BusServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RequestBus = channel.unary_unary(
                '/buses.BusService/RequestBus',
                request_serializer=buses__pb2.BusRequest.SerializeToString,
                response_deserializer=buses__pb2.BusResponse.FromString,
                )


class BusServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RequestBus(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BusServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RequestBus': grpc.unary_unary_rpc_method_handler(
                    servicer.RequestBus,
                    request_deserializer=buses__pb2.BusRequest.FromString,
                    response_serializer=buses__pb2.BusResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'buses.BusService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class BusService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RequestBus(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/buses.BusService/RequestBus',
            buses__pb2.BusRequest.SerializeToString,
            buses__pb2.BusResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
