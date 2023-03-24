package events

type HangupEvent struct {
	// Cause A numeric cause code for why the channel was hung up.
	Cause string
	// CauseTxt A description of why the channel was hung up.
	CauseTxt string
}

func (HangupEvent) EventTypeName() string {
	return "Hangup"
}
