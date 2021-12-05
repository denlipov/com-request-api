package api

import (
	"context"

	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
