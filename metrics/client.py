import grpc
import proto.metrics.pb.metrics_pb2_grpc as met_grpc, proto.metrics.pb.metrics_pb2 as met_pb

socket = 'localhost:50051'
channel = grpc.insecure_channel(socket)
client = met_grpc.MetricsServiceStub(channel)

