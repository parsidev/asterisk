package events

type Event interface {
	EventTypeName() string
}
