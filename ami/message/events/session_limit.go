package events

type SessionLimitEvent struct {
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

func (SessionLimitEvent) EventTypeName() string {
	return "SessionLimit"
}
