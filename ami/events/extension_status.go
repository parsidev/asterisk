package events

type ExtensionStatusEvent struct {
	Exten      string
	Context    string
	Hint       string
	Status     string
	StatusText string
}

func (ExtensionStatusEvent) EventTypeName() string {
	return "ExtensionStatus"
}
