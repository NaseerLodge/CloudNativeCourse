import grpc
import movieapi_pb2
import movieapi_pb2_grpc


def movieapi(st, n):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = movieapi_pb2_grpc.Movie_RecommendationStub(channel)
        request = movieapi_pb2.MovieRecommendation(moviename=st, no_of_recommendations=n)
        response = stub.Movie_Recommendation_by_cnn(request)
        return response.Recommended_Movies
