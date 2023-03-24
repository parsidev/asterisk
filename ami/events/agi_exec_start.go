package events

type AGIExecStartEvent struct {
	Command   string
	CommandId string
}

func (AGIExecStartEvent) EventTypeName() string {
	return "AGIExecStart"
}
