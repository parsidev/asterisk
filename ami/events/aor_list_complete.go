package events

type AorListCompleteEvent struct {
	EventList string
	ListItems string
}

func (AorListCompleteEvent) EventTypeName() string {
	return "AorListComplete"
}
