package sender

import (
	"errors"
	"log"
	"math/rand"

	"github.com/denlipov/com-request-api/internal/model"
)

type kafkaEventSender struct {
}

// NewEventSender ...
func NewEventSender() EventSender {
	return &kafkaEventSender{}
}

func (s *kafkaEventSender) Send(req *model.RequestEvent) error {
	ok := false
	if rand.Int63()%2 == 0 { // nolint:gosec
		ok = true
	}
	log.Printf("Send %v: %s", ok, req.String())

	if ok {
		return nil
	}
	return errors.New("Kafka sending failed")
}
