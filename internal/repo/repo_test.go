package repo

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/denlipov/com-request-api/internal/model"
	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func newMockRepo(t *testing.T) (*repo, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	repo := &repo{
		db: sqlxDB,
	}

	return repo, mock
}

type anyTime struct{}

func (a anyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestRemoveRequest_Success(t *testing.T) {
	r, dbMock := newMockRepo(t)
	ctx := context.Background()

	requestID := uint64(1)
	pbReq := &pb.Request{
		Id: requestID,
	}
	payload, err := protojson.Marshal(pbReq)
	if err != nil {
		t.Fatalf("Failed to marshal protobuf.Request: %+v", err)
	}

	dbMock.ExpectBegin()
	dbMock.ExpectExec("UPDATE requests").
		WithArgs(true, anyTime{}, requestID, false).
		WillReturnResult(sqlmock.NewResult(int64(requestID), 1))

	dbMock.ExpectExec("INSERT INTO requests_events").
		WithArgs(requestID,
			model.Removed,
			model.Idle,
			payload,
			anyTime{}).
		WillReturnResult(sqlmock.NewResult(int64(requestID), 1))
	dbMock.ExpectCommit()

	_, err = r.RemoveRequest(ctx, requestID)

	require.NoError(t, err)
}
