package events

type AorListEvent struct {
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
}

func (AorListEvent) EventTypeName() string {
	return "AorList"
}
