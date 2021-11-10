package repo

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/denlipov/com-request-api/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (r *repo) CreateRequest(ctx context.Context, req model.Request) (requestID uint64, err error) {

	doInsert := func(ctx context.Context, req model.Request, tx *sqlx.Tx) (uint64, error) {

		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(tx)

		// Request
		query := psql.
			Insert("requests").
			Columns(
				"service",
				"\"user\"",
				"text",
				"removed",
				"created",
				"updated").
			Values(
				req.Service,
				req.User,
				req.Text,
				false,
				time.Now(),
				time.Now()).
			Suffix("RETURNING id")

		var reqID uint64
		err := query.QueryRowContext(ctx).Scan(&reqID)
		if err != nil {
			return 0, err
		}

		// RequestEvent
		payload := fmt.Sprintf(
			`{
				"request": {
					"id": %d,
					"service": "%s",
					"user": "%s",
					"text": "%s"
				}
			}`,
			reqID, req.Service, req.User, req.Text)

		query = psql.
			Insert("requests_events").
			Columns(
				"request_id",
				"type",
				"locked",
				"payload",
				"updated").
			Values(
				reqID,
				"Created",
				false,
				payload,
				time.Now())

		_, err = query.ExecContext(ctx)
		if err != nil {
			return 0, err
		}

		return reqID, nil
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, errors.Wrap(err, "db.BeginTxx()")
	}

	reqID, err := doInsert(ctx, req, tx)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return 0, errors.Wrap(err, "Tx.Rollback()")
		}
		return 0, errors.Wrap(err, "Tx.WithTxFunc()")
	}

	if err = tx.Commit(); err != nil {
		return 0, errors.Wrap(err, "Tx.Commit()")
	}
	return reqID, nil
}
