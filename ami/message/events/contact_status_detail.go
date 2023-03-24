package events

type ContactStatusDetailEvent struct {
	AOR                 string
	URI                 string
	Status              string
	RoundtripUsec       string
	EndpointName        string
	UserAgent           string
	RegExpire           string
	ViaAddress          string
	CallID              string
	Id                  string
	AuthenticateQualify string
	OutboundProxy       string
	Path                string
	QualifyFrequency    string
	QualifyTimeout      string
}

func (ContactStatusDetailEvent) EventTypeName() string {
	return "ContactStatusDetail"
}
