package events

type UserEventEvent struct {
	UserEvent string
}

func (UserEventEvent) EventTypeName() string {
	return "UserEvent"
}
