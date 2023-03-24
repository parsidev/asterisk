package events

type ChallengeResponseFailedEvent struct {
	EventTV          string
	Severity         string
	Service          string
	EventVersion     string
	AccountID        string
	SessionID        string
	LocalAddress     string
	RemoteAddress    string
	Challenge        string
	Response         string
	ExpectedResponse string
	Module           string
	SessionTV        string
}

func (ChallengeResponseFailedEvent) EventTypeName() string {
	return "ChallengeResponseFailed"
}
