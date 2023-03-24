package events

type NewchannelEvent struct {
}

func (NewchannelEvent) EventTypeName() string {
	return "Newchannel"
}
