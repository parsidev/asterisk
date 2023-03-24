package events

type HangupHandlerRunEvent struct {
	Handler string
}

func (HangupHandlerRunEvent) EventTypeName() string {
	return "HangupHandlerRun"
}
