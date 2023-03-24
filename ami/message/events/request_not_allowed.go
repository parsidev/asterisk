package events

type RequestNotAllowedEvent struct {
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
	RequestParams string
}

func (RequestNotAllowedEvent) EventTypeName() string {
	return "RequestNotAllowed"
}
