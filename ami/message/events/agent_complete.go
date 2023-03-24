package events

type AgentCompleteEvent struct {
	Queue      string
	MemberName string
	Interface  string
	HoldTime   string
	TalkTime   string
	Reason     string
}

func (AgentCompleteEvent) EventTypeName() string {
	return "AgentComplete"
}
