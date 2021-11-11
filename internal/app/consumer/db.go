package consumer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/denlipov/com-request-api/internal/app/repo"
	"github.com/denlipov/com-request-api/internal/model"
)

// Consumer ...
type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	n      uint64
	events chan<- model.RequestEvent

	repo repo.EventRepo

	batchSize uint64
	timeout   time.Duration

	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

/*
type Config struct {
	n         uint64
	events    chan<- model.RequestEvent
	repo      repo.EventRepo
	batchSize uint64
	timeout   time.Duration
}
*/

// NewDbConsumer ...
func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.RequestEvent) Consumer {

	wg := &sync.WaitGroup{}

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        wg,
	}
}

func (c *consumer) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel

	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func(ctx context.Context) {

			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)

			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						continue
					}
					for _, event := range events {
						c.events <- event
					}
				case <-ctx.Done():
					log.Println("Consumer complete")
					return
				}
			}
		}(ctx)
	}
}

func (c *consumer) Close() {
	c.cancel()
	c.wg.Wait()
}
