package response

import "time"

// Model is a use case response model
type Model struct {
	Messages <-chan *Event
}

type Event struct {
	UserID    UserID
	SessionID SessionID
	MessageID MessageID
	Type      EvenType
	Time      time.Time
	Payload   EventPayload
}

type (
	UserID      string
	SessionID   string
	MessageID   int64
	MessageText string
)

type EvenType string

const (
	NewMessage    EvenType = "new"
	EditMessage   EvenType = "edit"
	DeleteMessage EvenType = "delete"
)

type EventPayload struct {
	New    *EventPayloadNew
	Edit   *EventPayloadEdit
	Delete *EventPayloadDelete
}

type (
	EventPayloadNew struct {
		MessageText MessageText
	}
	EventPayloadEdit struct {
		MessageText MessageText
	}
	EventPayloadDelete struct {
	}
)
