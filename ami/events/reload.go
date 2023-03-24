package events

type ReloadEvent struct {
	Module string
	Status string
}

func (ReloadEvent) EventTypeName() string {
	return "Reload"
}
