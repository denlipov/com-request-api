import asyncio

from grpclib.client import Channel

from ozonmp.com_request_api.v1.com_request_api_grpc import ComRequestApiServiceStub
from ozonmp.com_request_api.v1.com_request_api_pb2 import DescribeRequestV1Request, ListRequestV1Request, CreateRequestV1Request, RemoveRequestV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = ComRequestApiServiceStub(channel)

        # Create
        req = CreateRequestV1Request()
        req.request.service = "serv"
        req.request.user = "user"
        req.request.text = "text"
        reply = await client.CreateRequestV1(req)
        print(reply)
        
        # Describe
        req = DescribeRequestV1Request(request_id=1)
        reply = await client.DescribeRequestV1(req)
        print(reply)

        # List
        req = ListRequestV1Request()
        reply = await client.ListRequestV1(req)
        print(reply)

        # Remove
        req = RemoveRequestV1Request(request_id=1)
        reply = await client.RemoveRequestV1(req)
        print(reply)


if __name__ == '__main__':
    asyncio.run(main())
