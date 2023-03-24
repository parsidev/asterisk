package events

type MeetmeTalkingEvent struct {
	Meetme   string
	User     string
	Duration string
	Status   string
}

func (MeetmeTalkingEvent) EventTypeName() string {
	return "MeetmeTalking"
}
