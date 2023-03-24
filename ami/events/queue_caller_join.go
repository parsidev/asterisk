package events

type QueueCallerJoinEvent struct {
	Queue    string
	Position string
	Count    string
}

func (QueueCallerJoinEvent) EventTypeName() string {
	return "QueueCallerJoin"
}
