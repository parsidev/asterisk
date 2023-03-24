package events

type EndpointListEvent struct {
	ObjectType     string
	ObjectName     string
	Transport      string
	Aor            string
	Auths          string
	OutboundAuths  string
	DeviceState    string
	ActiveChannels string
}

func (EndpointListEvent) EventTypeName() string {
	return "EndpointList"
}
