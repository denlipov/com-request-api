package repo

import (
	"github.com/denlipov/com-request-api/internal/app/repo/postgres"
	"github.com/jmoiron/sqlx"
)

// NewEventRepo ...
func NewEventRepo(db *sqlx.DB) EventRepo {
	return postgres.NewEventRepo(db)
}
