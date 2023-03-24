package events

type QueueCallerLeaveEvent struct {
	Queue    string
	Count    string
	Position string
}

func (QueueCallerLeaveEvent) EventTypeName() string {
	return "QueueCallerLeave"
}
