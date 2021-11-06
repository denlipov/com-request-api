package model

import "fmt"

// Request - request entity.
//type Request struct {
//	ID  uint64 `db:"id"`
//	Foo uint64 `db:"foo"`
//}

type Request struct {
	ID      uint64 `json:"id,omitempty"`
	Service string `json:"service,omitempty"`
	User    string `json:"user,omitempty"`
	Text    string `json:"desc,omitempty"`
}

type RequestEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Request
}

var (
	evTypeStr = map[EventType]string{
		Created: "Created",
		Removed: "Removed",
		Updated: "Updated",
	}
	evStatusStr = map[EventStatus]string{
		Idle:      "Idle",
		Deferred:  "Deferred",
		Processed: "Processed",
	}
)

func (r Request) String() string {
	return fmt.Sprintf("id: %d; user: %s; desc: %s",
		r.ID, r.User, r.Text)
}

func (e RequestEvent) String() string {
	return fmt.Sprintf("RequestEvent { id: %d; type: %s; status: %s }",
		e.ID, evTypeStr[e.Type], evStatusStr[e.Status])
}
