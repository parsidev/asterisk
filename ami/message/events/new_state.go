package events

type NewstateEvent struct {
}

func (NewstateEvent) EventTypeName() string {
	return "Newstate"
}
