package api

import (
	"context"
	"fmt"
	"time"

	"github.com/denlipov/com-request-api/internal/model"
	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/go-redis/cache/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func requestKey(id uint64) string {
	return fmt.Sprintf("request:%d", id)
}

func (o *requestAPI) CreateRequestV1(
	ctx context.Context,
	req *pb.CreateRequestV1Request) (*pb.CreateRequestV1Response, error) {

	resp, err := o.createRequestV1(ctx, req)
	if err != nil {
		return nil, err
	}

	reqItem := &pb.Request{
		Id:      resp.RequestId,
		Service: req.Request.Service,
		User:    req.Request.User,
		Text:    req.Request.Text,
	}

	err = o.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   requestKey(resp.RequestId),
		Value: reqItem,
		TTL:   1 * time.Minute, // FIXME TTL
	})
	if err != nil {
		log.Error().Err(err).Msgf("unable to cache request: %s", reqItem.String())
	} else {
		log.Debug().Msgf("request cached: %+v", reqItem.String())
	}

	return resp, nil
}

func (o *requestAPI) createRequestV1(
	ctx context.Context,
	req *pb.CreateRequestV1Request) (*pb.CreateRequestV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.CreateRequestV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	requestID, err := o.repo.CreateRequest(ctx,
		model.Request{
			Service: req.Request.Service,
			User:    req.Request.User,
			Text:    req.Request.Text,
		})
	if err != nil {
		log.Error().Err(err).Msg("CreateRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateRequestV1 -- success")

	return &pb.CreateRequestV1Response{
		RequestId: requestID,
	}, nil
}
