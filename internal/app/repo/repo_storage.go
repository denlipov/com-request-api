package repo

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/denlipov/com-request-api/internal/model"
)

type memEventRepo struct {
	events map[uint64]*model.RequestEvent
	lock   sync.Mutex
}

// NewEventRepo ...
func NewEventRepo(storageCap uint64) EventRepo {
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
	return &memEventRepo{
		events: events,
	}
}

func (r *memEventRepo) Lock(n uint64) ([]model.RequestEvent, error) {
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
		log.Printf("Only %d events available; was requested: %d", eventsHoldTotal, n)
	}
	return result, nil
}

func (r *memEventRepo) Unlock(eventIDs []uint64) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	log.Printf("Unlocking events: %v", eventIDs)
	for _, id := range eventIDs {
		if _, exists := r.events[id]; exists {
			r.events[id].Status = model.Idle
		} else {
			log.Printf("Event ID %d does not exist", id)
		}
	}
	return nil
}

func (r *memEventRepo) Add(event []model.RequestEvent) error {
	return nil
}

func (r *memEventRepo) Remove(eventIDs []uint64) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	log.Printf("Removing events: %v", eventIDs)
	for _, id := range eventIDs {
		delete(r.events, id)
	}
	return nil
}
