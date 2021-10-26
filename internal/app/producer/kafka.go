package producer

import (
	"context"
	"log"
	"sync"
	"time"

	"com-request-api/internal/app/repo"
	"com-request-api/internal/app/sender"
	"com-request-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	repo   repo.EventRepo
	sender sender.EventSender
	events <-chan model.RequestEvent

	workerPool *workerpool.WorkerPool

	wg     *sync.WaitGroup
	cancel context.CancelFunc
}

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
							p.repo.Unlock(eventBuf)
						})
						eventsToUnlock = nil
					}

					if len(eventsToRemove) > 0 {
						eventBuf := eventsToRemove
						p.workerPool.Submit(func() {
							p.repo.Remove(eventBuf)
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
					log.Println("Producer complete")
					return
				}
			}
		}(ctx)
	}
}

func (p *producer) Close() {
	p.cancel()
	p.wg.Wait()
}
