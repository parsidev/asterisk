package events

type AlarmClearEvent struct {
	DAHDIChannel string
}

func (AlarmClearEvent) EventTypeName() string {
	return "AlarmClear"
}
