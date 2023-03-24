package events

type QueueMemberRinginuseEvent struct {
	Queue          string
	MemberName     string
	Interface      string
	StateInterface string
	Membership     string
	Penalty        string
	CallsTaken     string
	LastCall       string
	LastPause      string
	InCall         string
	Status         string
	Paused         string
	PausedReason   string
	Ringinuse      string
	Wrapuptime     string
}

func (QueueMemberRinginuseEvent) EventTypeName() string {
	return "QueueMemberRinginuse"
}
