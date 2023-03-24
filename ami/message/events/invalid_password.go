package events

type InvalidPasswordEvent struct {
	EventTV           string
	Severity          string
	Service           string
	EventVersion      string
	AccountID         string
	SessionID         string
	LocalAddress      string
	RemoteAddress     string
	Module            string
	SessionTV         string
	Challenge         string
	ReceivedChallenge string
	ReceivedHash      string
}

func (InvalidPasswordEvent) EventTypeName() string {
	return "InvalidPassword"
}
