package api

import (
	"context"
	"testing"

	"github.com/denlipov/com-request-api/internal/mocks"

	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/golang/mock/gomock"
)

func TestRequestAPI_Handlers(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepo(ctrl)
	reqAPI := NewRequestAPI(mockRepo)

	t.Run("DescribeRequestV1Request", func(t *testing.T) {
		invalidReqID := uint64(0)
		_, err := reqAPI.DescribeRequestV1(context.Background(),
			&pb.DescribeRequestV1Request{
				RequestId: invalidReqID,
			})
		if err == nil {
			t.Errorf("DescribeRequestV1Request() returned no error but expected error")
		}
	})

	t.Run("CreateRequestV1Request", func(t *testing.T) {

		testRequests := []*pb.Request{
			&pb.Request{
				Service: "x",
				User:    "user",
				Text:    "text",
			},
			&pb.Request{
				Service: "service",
				User:    "u",
				Text:    "text",
			},
			&pb.Request{
				Service: "service",
				User:    "user",
				Text:    "t",
			},
			&pb.Request{},
		}

		for _, testReq := range testRequests {
			_, err := reqAPI.CreateRequestV1(context.Background(), &pb.CreateRequestV1Request{
				Request: testReq,
			})
			if err == nil {
				t.Errorf("CreateRequestV1Request() returned no error but expected error")
			}
		}
	})

	t.Run("RemoveRequestV1Request", func(t *testing.T) {

		invalidReqID := uint64(0)
		_, err := reqAPI.RemoveRequestV1(context.Background(), &pb.RemoveRequestV1Request{
			RequestId: invalidReqID,
		})
		if err == nil {
			t.Errorf("RemoveRequestV1Request() returned no error but expected error")
		}
	})
}
