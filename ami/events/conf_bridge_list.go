package events

type ConfbridgeListEvent struct {
	Conference   string
	Admin        string
	MarkedUser   string
	WaitMarked   string
	EndMarked    string
	Waiting      string
	Muted        string
	Talking      string
	AnsweredTime string
}

func (ConfbridgeListEvent) EventTypeName() string {
	return "ConfbridgeList"
}
