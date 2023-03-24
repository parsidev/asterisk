package events

type BridgeLeaveEvent struct {
}

func (BridgeLeaveEvent) EventTypeName() string {
	return "BridgeLeave"
}
