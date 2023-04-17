from flask import Flask, request
import grpc
import stringmul_pb2
import stringmul_pb2_grpc

app = Flask(__name__)

@app.route("/multiply", methods=["POST"])
def multiply():
    st = request.form["string"]
    n = int(request.form["number"])
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = stringmul_pb2_grpc.StringMulStub(channel)
        request = stringmul_pb2.StringRequest(str=st, num=n)
        response = stub.Multiply(request)
        result = ", ".join(response.strings)
    return result

if __name__ == '__main__':
    app.run()
