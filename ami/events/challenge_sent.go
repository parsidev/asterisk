package events

type ChallengeSentEvent struct {
	EventTV       string
	Severity      string
	Service       string
	EventVersion  string
	AccountID     string
	SessionID     string
	LocalAddress  string
	RemoteAddress string
	Challenge     string
	Module        string
	SessionTV     string
}

func (ChallengeSentEvent) EventTypeName() string {
	return "ChallengeSent"
}
