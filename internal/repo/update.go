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

func (r *repo) UpdateRequest(ctx context.Context, req model.Request) (bool, error) {

	doUpdate := func(ctx context.Context, req model.Request, tx *sqlx.Tx) error {

		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(tx)

		payload := fmt.Sprintf(`{ "request": {"id": %d`, req.ID)
		// Request
		queryUpdate := psql.Update("requests")
		if req.Service != "" {
			queryUpdate = queryUpdate.Set("service", req.Service)
			payload += fmt.Sprintf(`, "service": "%s"`, req.Service)
		}
		if req.User != "" {
			queryUpdate = queryUpdate.Set("\"user\"", req.User)
			payload += fmt.Sprintf(`, "user": "%s"`, req.User)
		}
		if req.Text != "" {
			queryUpdate = queryUpdate.Set("text", req.Text)
			payload += fmt.Sprintf(`, "text": "%s"`, req.Text)
		}

		payload += "} }"

		queryUpdate = queryUpdate.Set("updated", time.Now()).
			Where(sq.And{
				sq.Eq{"id": req.ID},
				sq.Eq{"removed": false},
			})

		res, err := queryUpdate.ExecContext(ctx)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		} else if rowsAffected != 1 {
			return errors.New("No requests found to update or it was already removed")
		}

		// Event
		queryInsert := psql.
			Insert("requests_events").
			Columns(
				"request_id",
				"type",
				"locked",
				"payload",
				"updated").
			Values(
				req.ID,
				"Updated",
				false,
				payload,
				time.Now())

		_, err = queryInsert.ExecContext(ctx)
		if err != nil {
			return err
		}

		return nil
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return false, errors.Wrap(err, "db.BeginTxx()")
	}

	err = doUpdate(ctx, req, tx)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return false, errors.Wrap(err, "Tx.Rollback()")
		}
		return false, errors.Wrap(err, "Tx.WithTxFunc()")
	}

	if err = tx.Commit(); err != nil {
		return false, errors.Wrap(err, "Tx.Commit()")
	}
	return true, nil
}
