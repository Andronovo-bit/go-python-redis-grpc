import grpc
from concurrent import futures
import time

import user_pb2
import user_pb2_grpc

import Saver


class SaverServicer(user_pb2_grpc.SaverServicer):

    def SaveUser(self, request, context):
        response = user_pb2.UserReply()
        response.message = Saver.SaveUser(request)
        return response


server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

user_pb2_grpc.add_SaverServicer_to_server(SaverServicer(), server)

# listen on port 50051
print('Starting server. Listening on port 50051.')
server.add_insecure_port('[::]:50051')
server.start()

# since server.start() will not block,
# a sleep-loop is added to keep alive
try:
    while True:
        time.sleep(86400)
except KeyboardInterrupt:
    server.stop(0)
