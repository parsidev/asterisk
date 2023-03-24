package events

type AGIExecEndEvent struct {
	Command    string
	CommandId  string
	ResultCode string
	Result     string
}

func (AGIExecEndEvent) EventTypeName() string {
	return "AGIExecEnd"
}
