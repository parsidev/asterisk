package events

type ChanSpyStopEvent struct {
}

func (ChanSpyStopEvent) EventTypeName() string {
	return "ChanSpyStop"
}
