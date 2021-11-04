package sender

import (
	"github.com/denlipov/com-request-api/internal/model"
	"errors"
	"log"
	"math/rand"
)

type KafkaEventSender struct {
}

func NewEventSender() EventSender {
	return &KafkaEventSender{}
}

func (s *KafkaEventSender) Send(req *model.RequestEvent) error {
	ok := false
	if rand.Int63()%2 == 0 {
		ok = true
	}
	log.Printf("Send %v: %s", ok, req.String())

	if ok {
		return nil
	}
	return errors.New("Kafka sending failed")
}
