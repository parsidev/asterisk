package events

type MeetmeJoinEvent struct {
	Meetme string
	User   string
}

func (MeetmeJoinEvent) EventTypeName() string {
	return "MeetmeJoin"
}
