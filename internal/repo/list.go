package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/denlipov/com-request-api/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

func (r *repo) ListRequest(ctx context.Context, limit, offset uint64) ([]model.Request, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.ListRequest")
	defer span.Finish()

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "service", "\"user\"", "text").
		From("requests").
		Where(sq.And{
			sq.Gt{"id": offset},
			sq.Eq{"removed": false},
		}).
		OrderBy("id").
		Limit(limit).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	result := make([]model.Request, 0)
	err = sqlx.StructScan(rows, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
