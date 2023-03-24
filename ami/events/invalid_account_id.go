package events

type InvalidAccountIDEvent struct {
	EventTV       string
	Severity      string
	Service       string
	EventVersion  string
	AccountID     string
	SessionID     string
	LocalAddress  string
	RemoteAddress string
	Module        string
	SessionTV     string
}

func (InvalidAccountIDEvent) EventTypeName() string {
	return "InvalidAccountID"
}
