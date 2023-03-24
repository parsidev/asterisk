package events

type AgentLogoffEvent struct {
	Agent     string
	Logintime string
}

func (AgentLogoffEvent) EventTypeName() string {
	return "AgentLogoff"
}
