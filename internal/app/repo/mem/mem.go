package repo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/denlipov/com-request-api/internal/model"
)

// MemEventRepo ...
type MemEventRepo struct {
	events map[uint64]*model.RequestEvent
	lock   sync.Mutex
}

// NewEventRepo ...
func NewEventRepo(storageCap uint64) *MemEventRepo {
	events := make(map[uint64]*model.RequestEvent, storageCap)
	for i := uint64(0); i < storageCap; i++ {
		events[i] = &model.RequestEvent{
			ID:     i,
			Type:   model.Created,
			Status: model.Idle,
			Entity: &model.Request{
				ID:   uint64(rand.Int63()), // nolint:gosec
				User: "none",
				Text: fmt.Sprintf("req-%d", rand.Int63()), // nolint:gosec
			},
		}
	}
	return &MemEventRepo{
		events: events,
	}
}

// Lock ...
func (r *MemEventRepo) Lock(ctx context.Context, n uint64) ([]model.RequestEvent, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	result := make([]model.RequestEvent, n)
	eventsHoldTotal := uint64(0)
	for _, e := range r.events {
		if eventsHoldTotal == n {
			break
		}
		if e.Type == model.Created && e.Status == model.Idle {
			e.Status = model.Deferred
			result[eventsHoldTotal] = *e
			eventsHoldTotal++
		}
	}
	if eventsHoldTotal < n {
		result = result[:eventsHoldTotal]
		log.Debug().Msgf("Only %d events available; was requested: %d", eventsHoldTotal, n)
	}
	return result, nil
}

// Unlock ...
func (r *MemEventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	log.Debug().Msgf("Unlocking events: %v", eventIDs)
	for _, id := range eventIDs {
		if _, exists := r.events[id]; exists {
			r.events[id].Status = model.Idle
		} else {
			log.Debug().Msgf("Event ID %d does not exist", id)
		}
	}
	return nil
}

// Remove ...
func (r *MemEventRepo) Remove(ctx context.Context, eventIDs []uint64) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	log.Debug().Msgf("Removing events: %v", eventIDs)
	for _, id := range eventIDs {
		delete(r.events, id)
	}
	return nil
}
