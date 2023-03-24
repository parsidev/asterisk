package events

type DialStateEvent struct {
	DialStatus string
	Forward    string
}

func (DialStateEvent) EventTypeName() string {
	return "DialState"
}
