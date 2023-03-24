package events

type StatusCompleteEvent struct {
	Items string
}

func (StatusCompleteEvent) EventTypeName() string {
	return "StatusComplete"
}
