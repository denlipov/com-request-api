package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/denlipov/com-request-api/internal/model"
	"github.com/denlipov/com-request-api/internal/repo"
	"github.com/opentracing/opentracing-go"

	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
)

var (
	totalRequestNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "com_request_api_request_not_found_total",
		Help: "Total number of requests that were not found",
	})
)

type requestAPI struct {
	pb.UnimplementedComRequestApiServiceServer
	repo repo.Repo
}

// NewRequestAPI returns api of com-request-api service
func NewRequestAPI(r repo.Repo) pb.ComRequestApiServiceServer {
	return &requestAPI{repo: r}
}

func (o *requestAPI) DescribeRequestV1(
	ctx context.Context,
	req *pb.DescribeRequestV1Request,
) (*pb.DescribeRequestV1Response, error) {

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

	return &pb.DescribeRequestV1Response{
		Value: &pb.Request{
			Id:      request.ID,
			Service: request.Service,
			User:    request.User,
			Text:    request.Text,
		},
	}, nil
}

func (o *requestAPI) CreateRequestV1(
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

func (o *requestAPI) ListRequestV1(
	ctx context.Context,
	req *pb.ListRequestV1Request) (*pb.ListRequestV1Response, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "api.ListRequestV1")
	defer span.Finish()

	respArrayInternal, err := o.repo.ListRequest(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Error().Err(err).Msg("ListRequestV1 -- failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("ListRequestV1 -- success")

	respArray := make([]*pb.Request, len(respArrayInternal))
	for i, r := range respArrayInternal {
		respArray[i] = &pb.Request{
			Id:      r.ID,
			Service: r.Service,
			User:    r.User,
			Text:    r.Text,
		}
	}
	return &pb.ListRequestV1Response{
		Request: respArray,
	}, nil
}

func (o *requestAPI) RemoveRequestV1(
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

func (o *requestAPI) UpdateRequestV1(
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
