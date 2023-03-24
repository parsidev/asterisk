package events

type ChanSpyStartEvent struct {
}

func (ChanSpyStartEvent) EventTypeName() string {
	return "ChanSpyStart"
}
