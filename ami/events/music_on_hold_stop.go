package events

type MusicOnHoldStopEvent struct {
}

func (MusicOnHoldStopEvent) EventTypeName() string {
	return "MusicOnHoldStop"
}
