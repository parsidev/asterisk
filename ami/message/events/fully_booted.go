package events

type FullyBootedEvent struct {
	Status     string
	Uptime     string
	LastReload string
}

func (FullyBootedEvent) EventTypeName() string {
	return "FullyBooted"
}
