package repo

import (
	"context"

	"github.com/denlipov/com-request-api/internal/model"
)

type EventRepo interface {
	Lock(ctx context.Context, n uint64) ([]model.RequestEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error

	Remove(ctx context.Context, eventIDs []uint64) error
}
