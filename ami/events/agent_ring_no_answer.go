package events

type AgentRingNoAnswerEvent struct {
	Queue      string
	MemberName string
	Interface  string
	RingTime   string
}

func (AgentRingNoAnswerEvent) EventTypeName() string {
	return "AgentRingNoAnswer"
}
