package events

type BridgeCreateEvent struct {
}

func (BridgeCreateEvent) EventTypeName() string {
	return "BridgeCreate"
}
