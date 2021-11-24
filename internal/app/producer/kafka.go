package producer

import (
	"context"
	"sync"
	"time"

	"github.com/denlipov/com-request-api/internal/app/repo"
	"github.com/denlipov/com-request-api/internal/app/sender"
	"github.com/denlipov/com-request-api/internal/model"
	"github.com/rs/zerolog/log"

	"github.com/gammazero/workerpool"
)

// Producer ...
type Producer interface {
	Start()
	Close()
}

type producer struct {
	n uint64
	// timeout time.Duration

	repo   repo.EventRepo
	sender sender.EventSender
	events <-chan model.RequestEvent

	workerPool *workerpool.WorkerPool

	wg     *sync.WaitGroup
	cancel context.CancelFunc
}

// NewKafkaProducer ...
func NewKafkaProducer(
	n uint64,
	repo repo.EventRepo,
	sender sender.EventSender,
	events <-chan model.RequestEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}

	return &producer{
		n:          n,
		repo:       repo,
		sender:     sender,
		events:     events,
		workerPool: workerPool,
		wg:         wg,
	}
}

func (p *producer) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel

	log.Info().Msg("Producer started")
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func(ctx context.Context) {
			defer p.wg.Done()

			eventsToUnlock := make([]uint64, 0)
			eventsToRemove := make([]uint64, 0)

			ticker := time.NewTicker(2 * time.Second)

			for {
				select {
				case <-ticker.C:
					if len(eventsToUnlock) > 0 {
						eventBuf := eventsToUnlock
						p.workerPool.Submit(func() {
							err := p.repo.Unlock(ctx, eventBuf)
							if err != nil {
								log.Error().Err(err).Msg("Unlock() events failed")
							} else {
								repo.TotalEventsProcessedAdd(float64(len(eventsToUnlock)))
							}
						})
						eventsToUnlock = nil
					}

					if len(eventsToRemove) > 0 {
						eventBuf := eventsToRemove
						count := len(eventsToRemove)
						p.workerPool.Submit(func() {
							err := p.repo.Remove(ctx, eventBuf)
							if err != nil {
								log.Error().Err(err).Msg("Remove() events failed")
							} else {
								repo.TotalEventsProcessedAdd(float64(count))
							}
						})
						eventsToRemove = nil
					}

				case event := <-p.events:
					if err := p.sender.Send(&event); err != nil {
						eventsToUnlock = append(eventsToUnlock, event.ID)
					} else {
						eventsToRemove = append(eventsToRemove, event.ID)
					}

				case <-ctx.Done():
					log.Info().Msg("Producer complete")
					return
				}
			}
		}(ctx)
	}
}

func (p *producer) Close() {
	p.cancel()
	p.sender.Close()
	p.wg.Wait()
}
