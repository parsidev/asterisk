package events

type DTMFEndEvent struct {
	Digit      string
	DurationMs string
	Direction  string
}

func (DTMFEndEvent) EventTypeName() string {
	return "DTMFEnd"
}
