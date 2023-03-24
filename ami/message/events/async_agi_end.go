package events

type AsyncAGIEndEvent struct {
}

func (AsyncAGIEndEvent) EventTypeName() string {
	return "AsyncAGIEnd"
}
