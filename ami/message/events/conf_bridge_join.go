package events

type ConfbridgeJoinEvent struct {
	Conference string
	Admin      string
	Muted      string
}

func (ConfbridgeJoinEvent) EventTypeName() string {
	return "ConfbridgeJoin"
}
