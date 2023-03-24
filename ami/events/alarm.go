package events

type AlarmEvent struct {
	DAHDIChannel string
	Alarm        string
}

func (AlarmEvent) EventTypeName() string {
	return "Alarm"
}
