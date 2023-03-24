package events

type SuccessfulAuthEvent struct {
	EventTV       string
	Severity      string
	Service       string
	EventVersion  string
	AccountID     string
	SessionID     string
	LocalAddress  string
	RemoteAddress string
	UsingPassword string
	Module        string
	SessionTV     string
}

func (SuccessfulAuthEvent) EventTypeName() string {
	return "SuccessfulAuth"
}
