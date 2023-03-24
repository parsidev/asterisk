package events

type MusicOnHoldStartEvent struct {
	Class string
}

func (MusicOnHoldStartEvent) EventTypeName() string {
	return "MusicOnHoldStart"
}
