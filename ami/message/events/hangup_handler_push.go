package events

type HangupHandlerPushEvent struct {
	// Handler Hangup handler parameter string passed to the Gosub application.
	Handler string
}

func (HangupHandlerPushEvent) EventTypeName() string {
	return "HangupHandlerPush"
}
