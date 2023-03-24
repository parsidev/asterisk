package events

type AsyncAGIStartEvent struct {
	Env string
}

func (AsyncAGIStartEvent) EventTypeName() string {
	return "AsyncAGIStart"
}
