package events

type BridgeEnterEvent struct {
	SwapUniqueid string
}

func (BridgeEnterEvent) EventTypeName() string {
	return "BridgeEnter"
}
