package api

import (
	"context"
	"time"

	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/go-redis/cache/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *requestAPI) DescribeRequestV1(
	ctx context.Context,
	req *pb.DescribeRequestV1Request,
) (*pb.DescribeRequestV1Response, error) {

	reqItem := new(pb.Request)
	err := o.cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   requestKey(req.RequestId),
		Value: reqItem,
		TTL:   1 * time.Minute, // FIXME TTL
		Do: func(item *cache.Item) (interface{}, error) {
			log.Debug().Msgf("no request in cache (%d) -> check DB", req.RequestId)
			pbRequestPtr, err := o.describeRequestV1(ctx, req)
			if err != nil {
				return nil, err
			}
			return pbRequestPtr, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.DescribeRequestV1Response{
		Value: reqItem,
	}, nil
}

func (o *requestAPI) describeRequestV1(
	ctx context.Context,
	req *pb.DescribeRequestV1Request,
) (*pb.Request, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.DescribeRequestV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	request, err := o.repo.DescribeRequest(ctx, req.RequestId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if request == nil {
		log.Error().Uint64("requestId", req.RequestId).Msg("request not found")
		totalRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "request not found")
	}

	log.Debug().Msg("DescribeRequestV1 -- success")

	return &pb.Request{
		Id:      request.ID,
		Service: request.Service,
		User:    request.User,
		Text:    request.Text,
	}, nil
}
