package events

type CoreShowChannelsCompleteEvent struct {
	ActionID  string
	EventList string
	ListItems string
}

func (CoreShowChannelsCompleteEvent) EventTypeName() string {
	return "CoreShowChannelsComplete"
}
