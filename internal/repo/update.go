package repo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/denlipov/com-request-api/internal/model"
	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

func (r *repo) UpdateRequest(ctx context.Context, req model.Request) (bool, error) {

	doUpdate := func(ctx context.Context, req model.Request, tx *sqlx.Tx) error {

		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(tx)

		// Request
		queryUpdate := psql.Update("requests")
		if req.Service != "" {
			queryUpdate = queryUpdate.Set("service", req.Service)
		}
		if req.User != "" {
			queryUpdate = queryUpdate.Set("\"user\"", req.User)
		}
		if req.Text != "" {
			queryUpdate = queryUpdate.Set("text", req.Text)
		}

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
		pbReq := &pb.Request{
			Id:      req.ID,
			Service: req.Service,
			User:    req.User,
			Text:    req.Text,
		}
		payload, err := protojson.Marshal(pbReq)
		if err != nil {
			return err
		}

		queryInsert := psql.
			Insert("requests_events").
			Columns(
				"request_id",
				"type",
				"status",
				"payload",
				"updated").
			Values(
				req.ID,
				model.Updated,
				model.Idle,
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
