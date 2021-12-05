package api

import (
	"context"

	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *requestAPI) RemoveRequestV1(
	ctx context.Context,
	req *pb.RemoveRequestV1Request) (*pb.RemoveRequestV1Response, error) {

	resp, err := o.removeRequestV1(ctx, req)
	if err != nil {
		return nil, err
	}

	if !resp.Status {
		return resp, nil
	}

	err = o.cache.Delete(ctx, requestKey(req.RequestId))
	if err != nil {
		log.Error().Err(err).Msgf("unable to delete request from cache: %d", req.RequestId)
	} else {
		log.Debug().Msgf("request removed from cache: %d", req.RequestId)
	}

	return resp, nil
}

func (o *requestAPI) removeRequestV1(
	ctx context.Context,
	req *pb.RemoveRequestV1Request) (*pb.RemoveRequestV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.RemoveRequestV1")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ok, err := o.repo.RemoveRequest(ctx, req.RequestId)
	if err != nil {
		log.Error().Err(err).Msg("RemoveRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !ok {
		log.Error().Uint64("requestId", req.RequestId).Msg("request not found")
		totalRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "request not found")
	}

	log.Debug().Msg("RemoveRequestV1 -- success")

	return &pb.RemoveRequestV1Response{
		Status: true,
	}, nil
}
