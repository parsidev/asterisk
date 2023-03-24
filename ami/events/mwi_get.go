package events

type MWIGetEvent struct {
	ActionID    string
	Mailbox     string
	OldMessages string
	NewMessages string
}

func (MWIGetEvent) EventTypeName() string {
	return "MWIGet"
}
