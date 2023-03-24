package events

type LoadEvent struct {
	Module string
	Status string
}

func (LoadEvent) EventTypeName() string {
	return "Load"
}
