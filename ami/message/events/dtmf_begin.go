package events

type DTMFBeginEvent struct {
	Digit     string
	Direction string
}

func (DTMFBeginEvent) EventTypeName() string {
	return "DTMFBegin"
}
