package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/denlipov/com-request-api/internal/model"
)

// Repo is DAO for Request
type Repo interface {
	DescribeRequest(ctx context.Context, requestID uint64) (*model.Request, error)
	CreateRequest(ctx context.Context, req model.Request) (requestID uint64, err error)
	RemoveRequest(ctx context.Context, requestID uint64) (bool, error)
	ListRequest(ctx context.Context, limit, offset uint64) ([]model.Request, error)
	UpdateRequest(ctx context.Context, req model.Request) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{
		db:        db,
		batchSize: batchSize,
	}
}
