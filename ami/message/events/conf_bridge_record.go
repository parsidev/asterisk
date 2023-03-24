package events

type ConfbridgeRecordEvent struct {
	Conference string
}

func (ConfbridgeRecordEvent) EventTypeName() string {
	return "ConfbridgeRecord"
}
