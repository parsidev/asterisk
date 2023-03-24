package events

type BlindTransferEvent struct {
	Result     string
	IsExternal string
	Context    string
	Extension  string
}

func (BlindTransferEvent) EventTypeName() string {
	return "BlindTransfer"
}
