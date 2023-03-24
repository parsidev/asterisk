package events

type InvalidTransportEvent struct {
	EventTV            string
	Severity           string
	Service            string
	EventVersion       string
	AccountID          string
	SessionID          string
	LocalAddress       string
	RemoteAddress      string
	AttemptedTransport string
	Module             string
	SessionTV          string
}

func (InvalidTransportEvent) EventTypeName() string {
	return "InvalidTransport"
}
