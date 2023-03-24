package events

type AgentsEvent struct {
	Agent         string
	Name          string
	Status        string
	TalkingToChan string
	CallStarted   string
	LoggedInTime  string
	ActionID      string
}

func (AgentsEvent) EventTypeName() string {
	return "Agents"
}
