package events

type LogChannelEvent struct {
	Channel string
	Enabled string
}

func (LogChannelEvent) EventTypeName() string {
	return "LogChannel"
}
