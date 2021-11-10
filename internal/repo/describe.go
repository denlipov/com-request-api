package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/denlipov/com-request-api/internal/model"
)

func (r *repo) DescribeRequest(ctx context.Context, requestID uint64) (*model.Request, error) {

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "service", "\"user\"", "text").
		From("requests").
		Where(sq.Eq{"id": requestID}).ToSql()

	if err != nil {
		return nil, err
	}

	req := new(model.Request)
	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
