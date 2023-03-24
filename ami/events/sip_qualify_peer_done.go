package events

type SIPQualifyPeerDoneEvent struct {
	Peer     string
	ActionID string
}

func (SIPQualifyPeerDoneEvent) EventTypeName() string {
	return "SIPQualifyPeerDone"
}
