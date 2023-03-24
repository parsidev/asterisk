package events

type MeetmeMuteEvent struct {
	Meetme   string
	User     string
	Duration string
	Status   string
}

func (MeetmeMuteEvent) EventTypeName() string {
	return "MeetmeMute"
}
