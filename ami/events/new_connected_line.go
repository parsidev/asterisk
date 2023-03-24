package events

type NewConnectedLineEvent struct {
}

func (NewConnectedLineEvent) EventTypeName() string {
	return "NewConnectedLine"
}
