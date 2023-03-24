package events

type ShutdownEvent struct {
	Shutdown string
	Restart  string
}

func (ShutdownEvent) EventTypeName() string {
	return "Shutdown"
}
