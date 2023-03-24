package events

type ChannelTalkingStopEvent struct {
	Duration string
}

func (ChannelTalkingStopEvent) EventTypeName() string {
	return "ChannelTalkingStop"
}
