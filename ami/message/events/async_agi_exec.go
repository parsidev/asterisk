package events

type AsyncAGIExecEvent struct {
	CommandID string
	Result    string
}

func (AsyncAGIExecEvent) EventTypeName() string {
	return "AsyncAGIExec"
}
