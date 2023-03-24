package events

type UnloadEvent struct {
	Module string
	Status string
}

func (UnloadEvent) EventTypeName() string {
	return "Unload"
}
