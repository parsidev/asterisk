package events

type PeerStatusEvent struct {
	ChannelType string
	Peer        string
	PeerStatus  string
	Cause       string
	Address     string
	Port        string
	Time        float64
}

func (PeerStatusEvent) EventTypeName() string {
	return "PeerStatus"
}
