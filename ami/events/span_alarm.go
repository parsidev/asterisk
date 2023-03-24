package events

type SpanAlarmEvent struct {
	Span  string
	Alarm string
}

func (SpanAlarmEvent) EventTypeName() string {
	return "SpanAlarm"
}
