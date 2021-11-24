package retranslator

import (
	"testing"
	"time"

	"github.com/denlipov/com-request-api/internal/mocks"
	"github.com/denlipov/com-request-api/internal/model"
	"github.com/golang/mock/gomock"
)

func TestStart(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Retranslator test", func(t *testing.T) {

		batchSize := uint64(1)

		repo := mocks.NewMockEventRepo(ctrl)

		repo.EXPECT().
			Lock(gomock.Any(), batchSize).
			Return(
				[]model.RequestEvent{
					model.RequestEvent{ID: 1},
				},
				nil).
			AnyTimes()

		repo.EXPECT().
			Remove(gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()

		sender := mocks.NewMockEventSender(ctrl)

		sender.EXPECT().
			Send(&model.RequestEvent{ID: 1}).
			Return(nil).
			AnyTimes()

		sender.EXPECT().
			Close().
			AnyTimes()

		cfg := Config{
			ChannelSize:    32,
			ConsumerCount:  2,
			ConsumeSize:    batchSize,
			ConsumeTimeout: 1 * time.Second,
			ProducerCount:  2,
			WorkerCount:    2,
			Repo:           repo,
			Sender:         sender,
		}

		retranslator := NewRetranslator(cfg)
		retranslator.Start()
		time.Sleep(4 * time.Second)
		retranslator.Close()
	})
}
