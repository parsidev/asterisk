package events

type ConfbridgeEndEvent struct {
	Conference string
}

func (ConfbridgeEndEvent) EventTypeName() string {
	return "ConfbridgeEnd"
}
