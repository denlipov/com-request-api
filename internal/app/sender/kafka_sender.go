package sender

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/denlipov/com-request-api/internal/model"
	"github.com/rs/zerolog/log"
)

type kafkaEventSender struct {
}

// NewEventSender ...
func NewEventSender() EventSender {
	return &kafkaEventSender{}
}

func (s *kafkaEventSender) Send(ev *model.RequestEvent) error {
	ok := false
	if rand.Int63()%2 == 0 { // nolint:gosec
		ok = true
	}
	log.Info().Msg(fmt.Sprintf("Send %v: %s; req: %s", ok, ev.String(), ev.Entity.String()))

	if ok {
		return nil
	}
	return errors.New("Kafka sending failed")
}
