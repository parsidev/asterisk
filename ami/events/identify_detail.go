package events

type IdentifyDetailEvent struct {
	ObjectType   string
	ObjectName   string
	Endpoint     string
	SrvLookups   string
	Match        string
	MatchHeader  string
	EndpointName string
}

func (IdentifyDetailEvent) EventTypeName() string {
	return "IdentifyDetail"
}
