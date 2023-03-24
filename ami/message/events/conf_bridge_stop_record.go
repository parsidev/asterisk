package events

type ConfbridgeStopRecordEvent struct {
	Conference string
}

func (ConfbridgeStopRecordEvent) EventTypeName() string {
	return "ConfbridgeStopRecord"
}
