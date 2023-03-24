package events

type ConfbridgeStartEvent struct {
	Conference string
}

func (ConfbridgeStartEvent) EventTypeName() string {
	return "ConfbridgeStart"
}
