import asyncio

from grpclib.client import Channel

from ozonmp.com_request_api.v1.com_request_api_grpc import ComRequestApiServiceStub
from ozonmp.com_request_api.v1.com_request_api_pb2 import DescribeTemplateV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = ComRequestApiServiceStub(channel)

        req = DescribeRequestV1Request(template_id=1)
        reply = await client.DescribeRequestV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
