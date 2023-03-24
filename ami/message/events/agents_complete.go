package events

type AgentsCompleteEvent struct {
	ActionID string
}

func (AgentsCompleteEvent) EventTypeName() string {
	return "AgentsComplete"
}
