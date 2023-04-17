import grpc
from concurrent import futures
import stringmul_pb2
import stringmul_pb2_grpc

class StringMulServicer(stringmul_pb2_grpc.StringMulServicer):
    def Multiply(self, request, context):
        result = [request.str]
        for i in range(2, request.num + 1):
            result.append(request.str * i)
        return stringmul_pb2.StringArray(strings=result)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    stringmul_pb2_grpc.add_StringMulServicer_to_server(StringMulServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
