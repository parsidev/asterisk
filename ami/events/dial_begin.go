package events

type DialBeginEvent struct {
	DialString string
}

func (DialBeginEvent) EventTypeName() string {
	return "DialBegin"
}
