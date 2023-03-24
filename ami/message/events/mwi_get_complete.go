package events

type MWIGetCompleteEvent struct {
	ActionID  string
	EventList string
	ListItems string
}

func (MWIGetCompleteEvent) EventTypeName() string {
	return "MWIGetComplete"
}
