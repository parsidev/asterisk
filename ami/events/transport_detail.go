package events

type TransportDetailEvent struct {
	ObjectType               string
	ObjectName               string
	Protocol                 string
	Bind                     string
	AsycOperations           string
	CaListFile               string
	CaListPath               string
	CertFile                 string
	PrivKeyFile              string
	Password                 string
	ExternalSignalingAddress string
	ExternalSignalingPort    string
	ExternalMediaAddress     string
	Domain                   string
	VerifyServer             string
	VerifyClient             string
	RequireClientCert        string
	Method                   string
	Cipher                   string
	LocalNet                 string
	Tos                      string
	Cos                      string
	WebsocketWriteTimeout    string
	EndpointName             string
}

func (TransportDetailEvent) EventTypeName() string {
	return "TransportDetail"
}
