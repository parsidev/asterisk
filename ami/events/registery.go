package events

type RegistryEvent struct {
	ChannelType string
	UserName    string
	Domain      string
	Status      string
	Cause       string
}

func (RegistryEvent) EventTypeName() string {
	return "Registry"
}
