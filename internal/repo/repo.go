package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/denlipov/com-request-api/internal/model"
)

// Repo is DAO for Request
type Repo interface {
	DescribeRequest(ctx context.Context, requestID uint64) (*model.Request, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeRequest(ctx context.Context, requestID uint64) (*model.Request, error) {
	return nil, nil
}
