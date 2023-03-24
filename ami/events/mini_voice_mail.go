package events

type MiniVoiceMailEvent struct {
	Action  string
	Mailbox string
	Counter string
}

func (MiniVoiceMailEvent) EventTypeName() string {
	return "MiniVoiceMail"
}
