package events

type HangupHandlerPopEvent struct {
	// Handler Hangup handler parameter string passed to the Gosub application.
	Handler string
}

func (HangupHandlerPopEvent) EventTypeName() string {
	return "HangupHandlerPop"
}
