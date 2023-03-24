package events

type LocalBridgeEvent struct {
	Context           string
	Exten             string
	LocalOptimization string
}

func (LocalBridgeEvent) EventTypeName() string {
	return "LocalBridge"
}
