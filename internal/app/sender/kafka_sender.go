package sender

import (
	"context"
	"encoding/binary"
	"encoding/json"

	"github.com/denlipov/com-request-api/internal/model"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type kafkaEventSender struct {
	producer *kafka.Writer
	topic    string
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewEventSender ...
func NewEventSender(brokers []string, topic string) (EventSender, error) {

	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		RequiredAcks: kafka.RequireAll,
	}

	ctx, cancel := context.WithCancel(context.Background())

	kafkaSender := &kafkaEventSender{
		producer: w,
		topic:    topic,
		ctx:      ctx,
		cancel:   cancel,
	}

	return kafkaSender, nil
}

func (s *kafkaEventSender) Send(ev *model.RequestEvent) error {

	msgBytes, err := json.Marshal(*ev)
	if err != nil {
		log.Error().Err(err).Msgf("Unable to marshal msg to JSON")
		return err
	}

	key := make([]byte, 8)
	binary.LittleEndian.PutUint64(key, ev.ID)
	msg := kafka.Message{
		Key:   key,
		Value: msgBytes,
	}

	err = s.producer.WriteMessages(s.ctx, msg)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send event %d", ev.ID)
		return err
	}
	log.Debug().Msgf("event %d sent OK", ev.ID)
	return nil
}

func (s *kafkaEventSender) Close() {
	s.producer.Close()
	s.cancel()
}
