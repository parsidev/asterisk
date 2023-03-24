package events

type ContactListCompleteEvent struct {
	EventList string
	ListItems string
}

func (ContactListCompleteEvent) EventTypeName() string {
	return "ContactListComplete"
}
