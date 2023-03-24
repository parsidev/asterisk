package events

type AuthMethodNotAllowedEvent struct {
	EventTV       string
	Severity      string
	Service       string
	EventVersion  string
	AccountID     string
	SessionID     string
	LocalAddress  string
	RemoteAddress string
	AuthMethod    string
	Module        string
	SessionTV     string
}

func (AuthMethodNotAllowedEvent) EventTypeName() string {
	return "AuthMethodNotAllowed"
}
