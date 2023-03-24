package events

type ConfbridgeTalkingEvent struct {
	Conference    string
	TalkingStatus string
	Admin         string
}

func (ConfbridgeTalkingEvent) EventTypeName() string {
	return "ConfbridgeTalking"
}
