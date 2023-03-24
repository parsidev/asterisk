package events

type EndpointListCompleteEvent struct {
	EventList string
	ListItems string
}

func (EndpointListCompleteEvent) EventTypeName() string {
	return "EndpointListComplete"
}
