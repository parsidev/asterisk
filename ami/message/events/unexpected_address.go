package events

type UnexpectedAddressEvent struct {
	EventTV         string
	Severity        string
	Service         string
	EventVersion    string
	AccountID       string
	SessionID       string
	LocalAddress    string
	RemoteAddress   string
	ExpectedAddress string
	Module          string
	SessionTV       string
}

func (UnexpectedAddressEvent) EventTypeName() string {
	return "UnexpectedAddress"
}
