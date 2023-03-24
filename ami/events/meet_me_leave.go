package events

type MeetmeLeaveEvent struct {
	Meetme   string
	User     string
	Duration string
}

func (MeetmeLeaveEvent) EventTypeName() string {
	return "MeetmeLeave"
}
