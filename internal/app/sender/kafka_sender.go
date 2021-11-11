package sender

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/denlipov/com-request-api/internal/model"
	"github.com/rs/zerolog/log"
)

type KafkaEventSender struct {
}

func NewEventSender() EventSender {
	return &KafkaEventSender{}
}

func (s *KafkaEventSender) Send(ev *model.RequestEvent) error {
	ok := false
	if rand.Int63()%2 == 0 {
		ok = true
	}
	log.Info().Msg(fmt.Sprintf("Send %v: %s; req: %s", ok, ev.String(), ev.Entity.String()))

	if ok {
		return nil
	}
	return errors.New("Kafka sending failed")
}
