package repo

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/denlipov/com-request-api/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setupRepo(t *testing.T) (*repo, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	r := &repo{
		db: sqlxDB,
	}

	return r, mock
}

type anyTime struct{}

func (a anyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestRemoveRequest_Success(t *testing.T) {
	r, dbMock := setupRepo(t)
	ctx := context.Background()

	dbMock.ExpectBegin()
	dbMock.ExpectExec("UPDATE requests").
		WithArgs(true, anyTime{}, 1, false).
		WillReturnResult(sqlmock.NewResult(1, 1))

	dbMock.ExpectExec("INSERT INTO requests_events").
		WithArgs(1,
			model.Removed,
			model.Idle,
			[]byte{123, 34, 105, 100, 34, 58, 34, 49, 34, 125},
			anyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()

	_, err := r.RemoveRequest(ctx, 1)

	require.NoError(t, err)
}
