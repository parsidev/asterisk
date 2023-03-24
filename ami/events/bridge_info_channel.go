package events

type BridgeInfoChannelEvent struct {
}

func (BridgeInfoChannelEvent) EventTypeName() string {
	return "BridgeInfoChannel"
}
