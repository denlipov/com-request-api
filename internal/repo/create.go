package repo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/denlipov/com-request-api/internal/model"
	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

func (r *repo) CreateRequest(ctx context.Context, req model.Request) (uint64, error) {

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

		err := query.QueryRowContext(ctx).Scan(&req.ID)
		if err != nil {
			return 0, err
		}

		if span := opentracing.SpanFromContext(ctx); span != nil {
			span.SetTag("requestID", req.ID)
		}

		pbReq := &pb.Request{
			Id:      req.ID,
			Service: req.Service,
			User:    req.User,
			Text:    req.Text,
		}
		payload, err := protojson.Marshal(pbReq)
		if err != nil {
			return 0, err
		}

		// RequestEvent
		query = psql.
			Insert("requests_events").
			Columns(
				"request_id",
				"type",
				"status",
				"payload",
				"updated").
			Values(
				req.ID,
				model.Created,
				model.Idle,
				payload,
				time.Now())

		_, err = query.ExecContext(ctx)
		if err != nil {
			return 0, err
		}

		return req.ID, nil
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
