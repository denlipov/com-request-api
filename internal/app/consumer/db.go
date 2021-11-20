package consumer

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

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

	log.Info().Msg("Consumer started")
	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func(ctx context.Context) {

			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)

			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(ctx, c.batchSize)
					if err != nil {
						log.Error().Err(err).Msg("Lock() failed")
						continue
					}

					if len(events) > 0 {
						log.Debug().Msgf("%d events locked", len(events))
					}

					for _, event := range events {
						repo.TotalEventsProcessedAdd(1.0)
						c.events <- event
					}
				case <-ctx.Done():
					log.Info().Msg("Consumer complete")
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
