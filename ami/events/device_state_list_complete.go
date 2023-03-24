package events

type DeviceStateListCompleteEvent struct {
	EventList string
	ListItems string
}

func (DeviceStateListCompleteEvent) EventTypeName() string {
	return "DeviceStateListComplete"
}
