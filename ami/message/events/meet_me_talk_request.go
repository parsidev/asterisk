package events

type MeetmeTalkRequestEvent struct {
	Meetme   string
	User     string
	Duration string
	Status   string
}

func (MeetmeTalkRequestEvent) EventTypeName() string {
	return "MeetmeTalkRequest"
}
