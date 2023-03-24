package events

type AgentConnectEvent struct {
	Queue      string
	MemberName string
	Interface  string
	RingTime   string
	HoldTime   string
}

func (AgentConnectEvent) EventTypeName() string {
	return "AgentConnect"
}
