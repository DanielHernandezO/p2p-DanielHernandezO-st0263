# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import proto.search_file_pb2 as search__file__pb2


class SearchFileStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Search = channel.unary_unary(
                '/SearchFile/Search',
                request_serializer=search__file__pb2.FileRequested.SerializeToString,
                response_deserializer=search__file__pb2.Response.FromString,
                )


class SearchFileServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Search(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SearchFileServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Search': grpc.unary_unary_rpc_method_handler(
                    servicer.Search,
                    request_deserializer=search__file__pb2.FileRequested.FromString,
                    response_serializer=search__file__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'SearchFile', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SearchFile(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Search(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/SearchFile/Search',
            search__file__pb2.FileRequested.SerializeToString,
            search__file__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class JoinStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ExecuteJoin = channel.unary_unary(
                '/Join/ExecuteJoin',
                request_serializer=search__file__pb2.JoinPeersRequest.SerializeToString,
                response_deserializer=search__file__pb2.Response.FromString,
                )
        self.Join = channel.unary_unary(
                '/Join/Join',
                request_serializer=search__file__pb2.JoinRequest.SerializeToString,
                response_deserializer=search__file__pb2.Response.FromString,
                )


class JoinServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ExecuteJoin(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Join(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_JoinServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ExecuteJoin': grpc.unary_unary_rpc_method_handler(
                    servicer.ExecuteJoin,
                    request_deserializer=search__file__pb2.JoinPeersRequest.FromString,
                    response_serializer=search__file__pb2.Response.SerializeToString,
            ),
            'Join': grpc.unary_unary_rpc_method_handler(
                    servicer.Join,
                    request_deserializer=search__file__pb2.JoinRequest.FromString,
                    response_serializer=search__file__pb2.Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Join', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Join(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ExecuteJoin(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Join/ExecuteJoin',
            search__file__pb2.JoinPeersRequest.SerializeToString,
            search__file__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Join(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Join/Join',
            search__file__pb2.JoinRequest.SerializeToString,
            search__file__pb2.Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class FetchStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ExecuteFetch = channel.unary_unary(
                '/Fetch/ExecuteFetch',
                request_serializer=search__file__pb2.VoidRequest.SerializeToString,
                response_deserializer=search__file__pb2.Peers.FromString,
                )
        self.Fetch = channel.unary_unary(
                '/Fetch/Fetch',
                request_serializer=search__file__pb2.VoidRequest.SerializeToString,
                response_deserializer=search__file__pb2.Peers.FromString,
                )


class FetchServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ExecuteFetch(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Fetch(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_FetchServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ExecuteFetch': grpc.unary_unary_rpc_method_handler(
                    servicer.ExecuteFetch,
                    request_deserializer=search__file__pb2.VoidRequest.FromString,
                    response_serializer=search__file__pb2.Peers.SerializeToString,
            ),
            'Fetch': grpc.unary_unary_rpc_method_handler(
                    servicer.Fetch,
                    request_deserializer=search__file__pb2.VoidRequest.FromString,
                    response_serializer=search__file__pb2.Peers.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Fetch', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Fetch(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ExecuteFetch(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Fetch/ExecuteFetch',
            search__file__pb2.VoidRequest.SerializeToString,
            search__file__pb2.Peers.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Fetch(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Fetch/Fetch',
            search__file__pb2.VoidRequest.SerializeToString,
            search__file__pb2.Peers.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)