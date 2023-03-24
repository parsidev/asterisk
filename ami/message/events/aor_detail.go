package events

type AorDetailEvent struct {
	ObjectType          string
	ObjectName          string
	MinimumExpiration   string
	MaximumExpiration   string
	DefaultExpiration   string
	QualifyFrequency    string
	AuthenticateQualify string
	MaxContacts         string
	RemoveExisting      string
	Mailboxes           string
	OutboundProxy       string
	SupportPath         string
	TotalContacts       string
	ContactsRegistered  string
	EndpointName        string
}

func (AorDetailEvent) EventTypeName() string {
	return "AorDetail"
}
