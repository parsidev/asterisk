package events

type PresenceStatusEvent struct {
	Exten   string
	Context string
	Hint    string
	Status  string
	Subtype string
	Message string
}

func (PresenceStatusEvent) EventTypeName() string {
	return "PresenceStatus"
}
