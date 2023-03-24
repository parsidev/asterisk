package events

type MonitorStopEvent struct {
}

func (MonitorStopEvent) EventTypeName() string {
	return "MonitorStop"
}
