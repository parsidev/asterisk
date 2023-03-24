package events

type AuthListCompleteEvent struct {
	EventList string
	ListItems string
}

func (AuthListCompleteEvent) EventTypeName() string {
	return "AuthListComplete"
}
