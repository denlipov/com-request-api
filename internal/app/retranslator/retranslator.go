package retranslator

import (
	"time"

	"com-request-api/internal/app/consumer"
	"com-request-api/internal/app/producer"
	"com-request-api/internal/app/repo"
	"com-request-api/internal/app/sender"
	"com-request-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Retranslator interface {
	Start()
	Close()
}

type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.RequestEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

func fixConfig(c *Config) {
	const (
		DefaultRepoCapacity   = 1000
		DefaultConsumeTimeout = 2 * time.Second
	)

	if c.Repo == nil {
		c.Repo = repo.NewEventRepo(DefaultRepoCapacity)
	}
	if c.Sender == nil {
		c.Sender = sender.NewEventSender()
	}
	if c.ConsumeTimeout == 0 {
		c.ConsumeTimeout = DefaultConsumeTimeout
	}
}

func NewRetranslator(cfg Config) Retranslator {
	fixConfig(&cfg)
	events := make(chan model.RequestEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)

	producer := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.Repo,
		cfg.Sender,
		events,
		workerPool)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start() {
	r.producer.Start()
	r.consumer.Start()
}

func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
