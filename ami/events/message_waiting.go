package events

type MessageWaitingEvent struct {
	Mailbox string
	Waiting string
	New     string
	Old     string
}

func (MessageWaitingEvent) EventTypeName() string {
	return "MessageWaiting"
}
