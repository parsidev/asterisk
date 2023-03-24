package events

type PresenceStateChangeEvent struct {
	Presentity string
	Status     string
	Subtype    string
	Message    string
}

func (PresenceStateChangeEvent) EventTypeName() string {
	return "PresenceStateChange"
}
