package events

type BridgeInfoCompleteEvent struct {
}

func (BridgeInfoCompleteEvent) EventTypeName() string {
	return "BridgeInfoComplete"
}
