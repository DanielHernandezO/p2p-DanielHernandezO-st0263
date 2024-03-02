
import os
import grpc
import time
import proto.search_file_pb2 as search_file_pb2
import proto.search_file_pb2_grpc as search_file_pb2_grpc
from dotenv import load_dotenv
from datetime import datetime


def run():
    ip = os.getenv("ip")
    port = os.getenv("port")
    timeToSleep = os.getenv("time")
    channel = grpc.insecure_channel(f"{ip}:{port}")
    stub = search_file_pb2_grpc.FetchStub(channel)
    while True:
        response = stub.ExecuteFetch(search_file_pb2.VoidRequest()) 
        fecha_hora_actual = datetime.now().strftime("%Y-%m-%d %H:%M:%S") 
        with open('../logs/fetch.log', 'a') as archivo:
            archivo.write(f"{fecha_hora_actual}: {str(response.peers)}\n")
        time.sleep(int(timeToSleep))


if __name__ == "__main__":
    load_dotenv("configs/application.properties")
    run()
