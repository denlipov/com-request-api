package retranslator

import (
	"github.com/denlipov/com-request-api/internal/app/repo"
	"github.com/denlipov/com-request-api/internal/app/sender"
	"math/rand"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	t.Run("Retranslator test", func(t *testing.T) {
		rnd := func() uint64 {
			return uint64((rand.Int63() % 9) + int64(1))
		}

		cfg := Config{
			ChannelSize:    rnd() * 100,
			ConsumerCount:  rnd(),
			ConsumeSize:    rnd(),
			ConsumeTimeout: time.Duration(rnd()) * time.Second,
			ProducerCount:  rnd(),
			WorkerCount:    int(rnd()),
			Repo:           repo.NewEventRepo(rnd() * 1000),
			Sender:         sender.NewEventSender(),
		}

		retranslator := NewRetranslator(cfg)
		retranslator.Start()
		time.Sleep(time.Duration(rnd()) * time.Second)
		retranslator.Close()
	})
}
