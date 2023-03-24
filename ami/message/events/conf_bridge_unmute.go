package events

type ConfbridgeUnmuteEvent struct {
	Conference string
	Admin      string
}

func (ConfbridgeUnmuteEvent) EventTypeName() string {
	return "ConfbridgeUnmute"
}
