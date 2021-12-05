package api

import (
	"context"
	"time"

	"github.com/denlipov/com-request-api/internal/model"
	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/go-redis/cache/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *requestAPI) UpdateRequestV1(
	ctx context.Context,
	req *pb.UpdateRequestV1Request) (*pb.UpdateRequestV1Response, error) {

	resp, err := o.updateRequestV1(ctx, req)
	if err != nil {
		return nil, err
	}

	if !resp.Status {
		return resp, nil
	}

	reqItem := &pb.Request{
		Id:      req.RequestId,
		Service: req.Body.Service,
		User:    req.Body.User,
		Text:    req.Body.Text,
	}

	err = o.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   requestKey(req.RequestId),
		Value: reqItem,
		TTL:   1 * time.Minute, // FIXME TTL
	})
	if err != nil {
		log.Error().Err(err).Msgf("unable to update cache for request: %+v", reqItem.String())
	} else {
		log.Debug().Msgf("request cache updated: %+v", reqItem.String())
	}

	return resp, nil
}

func (o *requestAPI) updateRequestV1(
	ctx context.Context,
	req *pb.UpdateRequestV1Request) (*pb.UpdateRequestV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.UpdateRequestV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("UpdateRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mreq := model.Request{
		ID:      req.RequestId,
		Service: req.Body.Service,
		User:    req.Body.User,
		Text:    req.Body.Text,
	}
	ok, err := o.repo.UpdateRequest(ctx, mreq)
	if err != nil {
		log.Error().Err(err).Msg("UpdateRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !ok {
		log.Debug().Uint64("requestId", mreq.ID).Msg("request not found")
		totalRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "request not found")
	}

	log.Debug().Msg("UpdateRequestV1 -- success")

	return &pb.UpdateRequestV1Response{
		Status: true,
	}, nil
}
