package events

type FAXSessionsEntryEvent struct {
	ActionID      string
	Channel       string
	Technology    string
	SessionNumber string
	SessionType   string
	Operation     string
	State         string
	Files         string
}

func (FAXSessionsEntryEvent) EventTypeName() string {
	return "FAXSessionsEntry"
}
