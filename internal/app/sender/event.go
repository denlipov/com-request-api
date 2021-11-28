package sender

import (
	"github.com/denlipov/com-request-api/internal/model"
)

// EventSender ...
type EventSender interface {
	Send(event *model.RequestEvent) error
	Close()
}
