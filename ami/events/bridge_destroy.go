package events

type BridgeDestroyEvent struct {
}

func (BridgeDestroyEvent) EventTypeName() string {
	return "BridgeDestroy"
}
