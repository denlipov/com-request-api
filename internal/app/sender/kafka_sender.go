package sender

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/denlipov/com-request-api/internal/model"
	"github.com/rs/zerolog/log"
)

type kafkaEventSender struct {
	producer sarama.SyncProducer
	topic    string
	wg       *sync.WaitGroup
	cancel   context.CancelFunc
}

// NewEventSender ...
func NewEventSender(brokers []string, topic string) (EventSender, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := producer.Close(); err != nil {
			log.Error().Err(err).Msgf("Failed to close producer correctly")
		}
	}()

	kafkaSender := &kafkaEventSender{
		producer: producer,
		topic:    topic,
		wg:       wg,
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

	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(msgBytes),
	}

	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send event %d", ev.ID)
		return err
	}
	log.Debug().Msgf("event %d sent OK; part: %d; offset: %d", ev.ID, partition, offset)
	return nil
}

func (s *kafkaEventSender) Close() {
	s.cancel()
	s.wg.Wait()
}
