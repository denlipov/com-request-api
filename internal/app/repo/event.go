package repo

import (
	"github.com/denlipov/com-request-api/internal/model"
)

// EventRepo ...
type EventRepo interface {
	Lock(n uint64) ([]model.RequestEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.RequestEvent) error
	Remove(eventIDs []uint64) error
}
