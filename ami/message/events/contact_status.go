package events

type ContactStatusEvent struct {
	URI           string
	ContactStatus string
	AOR           string
	EndpointName  string
	RoundtripUsec string
}

func (ContactStatusEvent) EventTypeName() string {
	return "ContactStatus"
}
