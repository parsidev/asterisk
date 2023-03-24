package events

type MeetmeEndEvent struct {
	Meetme string
}

func (MeetmeEndEvent) EventTypeName() string {
	return "MeetmeEnd"
}
