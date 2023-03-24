package events

type HangupRequestEvent struct {
	Cause string
}

func (HangupRequestEvent) EventTypeName() string {
	return "HangupRequest"
}
