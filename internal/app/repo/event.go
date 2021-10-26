package repo

import (
	"com-request-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.RequestEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.RequestEvent) error
	Remove(eventIDs []uint64) error
}
