# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: ozonmp/com_request_api/v1/com_request_api.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import validate.validate_pb2
import google.api.annotations_pb2
import google.protobuf.timestamp_pb2
import ozonmp.com_request_api.v1.com_request_api_pb2


class ComRequestApiServiceBase(abc.ABC):

    @abc.abstractmethod
    async def DescribeRequestV1(self, stream: 'grpclib.server.Stream[ozonmp.com_request_api.v1.com_request_api_pb2.DescribeRequestV1Request, ozonmp.com_request_api.v1.com_request_api_pb2.DescribeRequestV1Response]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/denlipov.com_request_api.v1.ComRequestApiService/DescribeRequestV1': grpclib.const.Handler(
                self.DescribeRequestV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.com_request_api.v1.com_request_api_pb2.DescribeRequestV1Request,
                ozonmp.com_request_api.v1.com_request_api_pb2.DescribeRequestV1Response,
            ),
        }


class ComRequestApiServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.DescribeRequestV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/denlipov.com_request_api.v1.ComRequestApiService/DescribeRequestV1',
            ozonmp.com_request_api.v1.com_request_api_pb2.DescribeRequestV1Request,
            ozonmp.com_request_api.v1.com_request_api_pb2.DescribeRequestV1Response,
        )
