package events

type AttendedTransferEvent struct {
	Result                string
	DestType              string
	DestBridgeUniqueid    string
	DestApp               string
	DestTransfererChannel string
}

func (AttendedTransferEvent) EventTypeName() string {
	return "AttendedTransfer"
}
