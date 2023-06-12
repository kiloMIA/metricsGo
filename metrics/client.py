import grpc
import import proto.metrics.pb.metrics_pb2_grpc as met_grpc, import proto.metrics.pb.metrics_pb2 as met_pb
logging.basicConfig(format='%(asctime)s:%(name)s : %(message)s', level=logging.DEBUG)
log = logging.getLogger(__name__)
socket = 'localhost:50051'
channel = grpc.insecure_channel(socket)
client = met_grpc.MetricsServiceStub(channel)
log.info(f"Connected to the Server : {socket}")
