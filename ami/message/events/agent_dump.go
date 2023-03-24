package events

type AgentDumpEvent struct {
	Queue      string
	MemberName string
	Interface  string
}

func (AgentDumpEvent) EventTypeName() string {
	return "AgentDump"
}
