package events

type UnholdEvent struct {
}

func (UnholdEvent) EventTypeName() string {
	return "Unhold"
}
