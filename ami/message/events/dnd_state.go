package events

type DNDStateEvent struct {
	DAHDIChannel string
	Status       string
}

func (DNDStateEvent) EventTypeName() string {
	return "DNDState"
}
