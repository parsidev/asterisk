package events

type AgentCalledEvent struct {
	Queue      string
	MemberName string
	Interface  string
}

func (AgentCalledEvent) EventTypeName() string {
	return "AgentCalled"
}
