package events

type FAXSessionsCompleteEvent struct {
	ActionID string
	Total    string
}

func (FAXSessionsCompleteEvent) EventTypeName() string {
	return "FAXSessionsComplete"
}
