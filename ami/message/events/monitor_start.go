package events

type MonitorStartEvent struct {
}

func (MonitorStartEvent) EventTypeName() string {
	return "MonitorStart"
}
