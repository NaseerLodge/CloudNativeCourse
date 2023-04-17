import grpc
import stringmul_pb2
import stringmul_pb2_grpc

def stringmul(st, n):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = stringmul_pb2_grpc.StringMulStub(channel)
        request = stringmul_pb2.StringRequest(str=st, num=n)
        response = stub.Multiply(request)
        return response.strings

if __name__ == '__main__':
    print(stringmul('apple', 3))