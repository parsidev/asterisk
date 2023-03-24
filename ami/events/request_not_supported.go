package events

type RequestNotSupportedEvent struct {
	EventTV       string
	Severity      string
	Service       string
	EventVersion  string
	AccountID     string
	SessionID     string
	LocalAddress  string
	RemoteAddress string
	RequestType   string
	Module        string
	SessionTV     string
}

func (RequestNotSupportedEvent) EventTypeName() string {
	return "RequestNotSupported"
}
