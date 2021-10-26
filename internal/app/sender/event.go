package sender

import (
	"com-request-api/internal/model"
)

type EventSender interface {
	Send(event *model.RequestEvent) error
}
