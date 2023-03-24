package events

type NewAccountCodeEvent struct {
	OldAccountCode string
}

func (NewAccountCodeEvent) EventTypeName() string {
	return "NewAccountCode"
}
