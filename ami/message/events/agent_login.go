package events

type AgentLoginEvent struct {
	Agent string
}

func (AgentLoginEvent) EventTypeName() string {
	return "AgentLogin"
}
