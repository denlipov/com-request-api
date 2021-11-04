package sender

import (
	"github.com/denlipov/com-request-api/internal/model"
)

type EventSender interface {
	Send(event *model.RequestEvent) error
}
