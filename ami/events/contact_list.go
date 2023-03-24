package events

type ContactListEvent struct {
	// ObjectType The object's type. This will always be 'contact'.
	ObjectType string
	// ObjectName The name of this object.
	ObjectName string
	// ViaAddr IP address of the last Via header in REGISTER request. Will only appear in the event if available.
	ViaAddr string
	// ViaPort Port number of the last Via header in REGISTER request. Will only appear in the event if available.
	ViaPort string
	// QualifyTimeout The elapsed time in decimal seconds after which an OPTIONS message is sent before the contact is considered unavailable.
	QualifyTimeout string
	// CallId Content of the Call-ID header in REGISTER request. Will only appear in the event if available.
	CallId string
	// RegServer Asterisk Server name.
	RegServer string
	// PruneOnBoot If true delete the contact on Asterisk restart/boot.
	PruneOnBoot string
	// Path The Path header received on the REGISTER.
	Path string
	// Endpoint The name of the endpoint associated with this information.
	Endpoint string
	// AuthenticateQualify A boolean indicating whether a qualify should be authenticated.
	AuthenticateQualify string
	// Uri This contact's URI.
	Uri string
	// QualifyFrequency The interval in seconds at which the contact will be qualified.
	QualifyFrequency string
	// UserAgent Content of the User-Agent header in REGISTER request
	UserAgent string
	// ExpirationTime Absolute time that this contact is no longer valid after
	ExpirationTime string
	// OutboundProxy The contact's outbound proxy.
	OutboundProxy string
	// Status This contact's status.
	Status string
	// RoundtripUsec The round trip time in microseconds.
	RoundtripUsec string
}

func (ContactListEvent) EventTypeName() string {
	return "ContactList"
}
