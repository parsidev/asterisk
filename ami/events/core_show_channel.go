package events

type CoreShowChannelEvent struct {
	ActionID        string
	BridgeId        string
	Application     string
	ApplicationData string
	Duration        string
}

func (CoreShowChannelEvent) EventTypeName() string {
	return "CoreShowChannel"
}
