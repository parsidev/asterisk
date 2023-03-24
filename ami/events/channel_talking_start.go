package events

type ChannelTalkingStartEvent struct {
}

func (ChannelTalkingStartEvent) EventTypeName() string {
	return "ChannelTalkingStart"
}
