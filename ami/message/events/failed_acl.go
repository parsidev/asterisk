package events

type FailedACLEvent struct {
	EventTV       string
	Severity      string
	Service       string
	EventVersion  string
	AccountID     string
	SessionID     string
	LocalAddress  string
	RemoteAddress string
	Module        string
	ACLName       string
	SessionTV     string
}

func (FailedACLEvent) EventTypeName() string {
	return "FailedACL"
}
