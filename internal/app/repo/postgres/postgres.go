package postgres

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/denlipov/com-request-api/internal/model"
	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
)

// PGEventRepo ...
type PGEventRepo struct {
	db *sqlx.DB
}

// NewEventRepo ...
func NewEventRepo(db *sqlx.DB) *PGEventRepo {
	return &PGEventRepo{
		db: db,
	}
}

// Lock ...
func (r *PGEventRepo) Lock(ctx context.Context, n uint64) ([]model.RequestEvent, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.Lock")
	defer span.Finish()

	doLockTx := func(ctx context.Context, n uint64, tx *sqlx.Tx) ([]model.RequestEvent, error) {
		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

		querySelect, args, err := psql.
			Select("id", "type", "payload").
			From("requests_events").
			Where(sq.Eq{"status": model.Idle}).
			OrderBy("updated").
			Limit(n).
			ToSql()
		if err != nil {
			return nil, err
		}

		rows, err := tx.QueryxContext(ctx, querySelect, args...)
		if err != nil {
			return nil, err
		}

		evIDs := make([]uint64, 0)
		result := make([]model.RequestEvent, 0)
		for rows.Next() {
			var (
				id        uint64
				eventType model.EventType
				payload   []byte
				pbReq     pb.Request
			)

			err = rows.Scan(&id, &eventType, &payload)
			if err != nil {
				return nil, err
			}

			err = protojson.Unmarshal(payload, &pbReq)
			if err != nil {
				return nil, err
			}
			e := model.RequestEvent{
				ID:     id,
				Type:   eventType,
				Status: model.Deferred,
				Entity: &model.Request{
					ID:      pbReq.Id,
					Service: pbReq.Service,
					User:    pbReq.User,
					Text:    pbReq.Text,
				},
			}
			result = append(result, e)
			evIDs = append(evIDs, id)
		}

		if len(evIDs) == 0 {
			return nil, nil
		}

		// Update
		queryUpdate := psql.
			Update("requests_events").
			Set("status", model.Deferred).
			Set("updated", time.Now()).
			Where(sq.Eq{"id": evIDs}).
			RunWith(tx)

		_, err = queryUpdate.ExecContext(ctx)
		if err != nil {
			return nil, err
		}

		log.Debug().Msgf("Lock(): Updated records: %v", result)

		return result, nil
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "db.BeginTxx()")
	}

	result, err := doLockTx(ctx, n, tx)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return nil, errors.Wrap(err, "Tx.Rollback()")
		}
		return nil, errors.Wrap(err, "Tx.WithTxFunc()")
	}

	if err = tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "Tx.Commit()")
	}

	return result, nil
}

// Unlock ...
func (r *PGEventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {

	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.Unlock")
	defer span.Finish()

	if len(eventIDs) == 0 {
		return nil
	}

	log.Debug().Msgf("Unlock() called for records: %v", eventIDs)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.db)

	query := psql.Update("requests_events").
		Set("status", model.Idle).
		Set("updated", time.Now()).
		Where(sq.Eq{"id": eventIDs})

	res, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if int(rowsAffected) != len(eventIDs) {
		log.Info().Msgf("Unlock(): Updated only %d records; requested: %d", rowsAffected, len(eventIDs))
	}

	return nil
}

// Remove ...
func (r *PGEventRepo) Remove(ctx context.Context, eventIDs []uint64) error {

	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.Remove")
	defer span.Finish()

	if len(eventIDs) == 0 {
		return nil
	}

	log.Debug().Msgf("Remove() called for events: %v", eventIDs)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(r.db)
	query := psql.Delete("requests_events").
		Where(sq.Eq{"id": eventIDs})

	res, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if int(rowsAffected) != len(eventIDs) {
		log.Debug().Msgf("Remove(): Removed only %d records; requested: %d", rowsAffected, len(eventIDs))
	}

	return nil
}
