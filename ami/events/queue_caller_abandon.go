package events

type QueueCallerAbandonEvent struct {
	Queue            string
	Position         string
	OriginalPosition string
	HoldTime         string
}

func (QueueCallerAbandonEvent) EventTypeName() string {
	return "QueueCallerAbandon"
}
