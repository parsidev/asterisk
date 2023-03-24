package events

type NewCalleridEvent struct {
	CIDCallingPres string
}

func (NewCalleridEvent) EventTypeName() string {
	return "NewCallerid"
}
