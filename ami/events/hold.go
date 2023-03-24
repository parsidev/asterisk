package events

type HoldEvent struct {
	MusicClass string
}

func (HoldEvent) EventTypeName() string {
	return "Hold"
}
