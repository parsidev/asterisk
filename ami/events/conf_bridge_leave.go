package events

type ConfbridgeLeaveEvent struct {
	Conference string
	Admin      string
}

func (ConfbridgeLeaveEvent) EventTypeName() string {
	return "ConfbridgeLeave"
}
