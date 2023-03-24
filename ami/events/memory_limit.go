package events

type MemoryLimitEvent struct {
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

func (MemoryLimitEvent) EventTypeName() string {
	return "MemoryLimit"
}
