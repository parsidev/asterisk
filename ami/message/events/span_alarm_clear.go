package events

type SpanAlarmClearEvent struct {
	Span string
}

func (SpanAlarmClearEvent) EventTypeName() string {
	return "SpanAlarmClear"
}
