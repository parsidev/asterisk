package events

type FAXStatsEvent struct {
	ActionID         string
	CurrentSessions  string
	ReservedSessions string
	TransmitAttempts string
	ReceiveAttempts  string
	CompletedFAXes   string
	FailedFAXes      string
}

func (FAXStatsEvent) EventTypeName() string {
	return "FAXStats"
}
