package events

type SessionTimeoutEvent struct {
	Source string
}

func (SessionTimeoutEvent) EventTypeName() string {
	return "SessionTimeout"
}
