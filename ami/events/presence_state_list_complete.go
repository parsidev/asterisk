package events

type PresenceStateListCompleteEvent struct {
	EventList string
	ListItems string
}

func (PresenceStateListCompleteEvent) EventTypeName() string {
	return "PresenceStateListComplete"
}
