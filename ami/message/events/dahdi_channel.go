package events

type DAHDIChannelEvent struct {
	DAHDIGroup   string
	DAHDISpan    string
	DAHDIChannel string
}

func (DAHDIChannelEvent) EventTypeName() string {
	return "DAHDIChannel"
}
