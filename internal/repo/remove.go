package repo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

func (r *repo) RemoveRequest(ctx context.Context, requestID uint64) (bool, error) {

	doRemove := func(ctx context.Context, requestID uint64, tx *sqlx.Tx) error {

		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(tx)

		// Request
		queryUpdate := psql.
			Update("requests").
			Set("removed", true).
			Set("updated", time.Now()).
			Where(sq.And{
				sq.Eq{"id": requestID},
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

		pbReq := &pb.Request{
			Id: requestID,
		}
		payload, err := protojson.Marshal(pbReq)
		if err != nil {
			return err
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
				requestID,
				"Removed",
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

	err = doRemove(ctx, requestID, tx)
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
