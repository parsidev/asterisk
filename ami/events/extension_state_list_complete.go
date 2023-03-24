package events

type ExtensionStateListCompleteEvent struct {
	EventList string
	ListItems string
}

func (ExtensionStateListCompleteEvent) EventTypeName() string {
	return "ExtensionStateListComplete"
}
