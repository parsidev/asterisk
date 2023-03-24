package events

type DialEndEvent struct {
	DialStatus string
	Forward    string
}

func (DialEndEvent) EventTypeName() string {
	return "DialEnd"
}
