package events

type BridgeVideoSourceUpdateEvent struct {
	BridgePreviousVideoSource string
}

func (BridgeVideoSourceUpdateEvent) EventTypeName() string {
	return "BridgeVideoSourceUpdate"
}
