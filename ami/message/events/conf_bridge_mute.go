package events

type ConfbridgeMuteEvent struct {
	Conference string
	Admin      string
}

func (ConfbridgeMuteEvent) EventTypeName() string {
	return "ConfbridgeMute"
}
