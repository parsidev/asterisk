package events

type RenameEvent struct {
	Channel  string
	NewName  string
	Uniqueid string
}

func (RenameEvent) EventTypeName() string {
	return "Rename"
}
