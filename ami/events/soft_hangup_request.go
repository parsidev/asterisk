package events

type SoftHangupRequestEvent struct {
	Cause string
}

func (SoftHangupRequestEvent) EventTypeName() string {
	return "SoftHangupRequest"
}
