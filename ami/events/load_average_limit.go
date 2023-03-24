package events

type LoadAverageLimitEvent struct {
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

func (LoadAverageLimitEvent) EventTypeName() string {
	return "LoadAverageLimit"
}
