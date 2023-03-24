package events

type EndpointDetailCompleteEvent struct {
	EventList string
	ListItems string
}

func (EndpointDetailCompleteEvent) EventTypeName() string {
	return "EndpointDetailComplete"
}
