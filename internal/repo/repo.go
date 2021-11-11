package repo

import (
	"context"
	"math/rand"

	"github.com/jmoiron/sqlx"

	"github.com/denlipov/com-request-api/internal/model"
)

// Repo is DAO for Request
type Repo interface {
	DescribeRequest(ctx context.Context, requestID uint64) (*model.Request, error)
	CreateRequest(ctx context.Context, req model.Request) (requestID uint64, err error)
	RemoveRequest(ctx context.Context, requestID uint64) (bool, error)
	ListRequest(ctx context.Context) ([]model.Request, error)
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
	return &model.Request{
		ID:      requestID,
		Service: "dummyService",
		User:    "dummyUser",
		Text:    "dummyText",
	}, nil
}

func (r *repo) CreateRequest(ctx context.Context, req model.Request) (requestID uint64, err error) {
	return uint64(rand.Int63()), nil // nolint:gosec
}

func (r *repo) RemoveRequest(ctx context.Context, requestID uint64) (bool, error) {
	return true, nil
}

func (r *repo) ListRequest(ctx context.Context) ([]model.Request, error) {
	return []model.Request{
		{
			ID:      uint64(rand.Int63()), // nolint:gosec
			Service: "someService1",
			User:    "someUser1",
			Text:    "someText1",
		},
		{
			ID:      uint64(rand.Int63()), // nolint:gosec
			Service: "someService2",
			User:    "someUser2",
			Text:    "someText2",
		},
	}, nil
}
