package events

type SendFAXEvent struct {
	// LocalStationID The value of the `LOCALSTATIONID` channel variable
	LocalStationID string
	// RemoteStationID The value of the `REMOTESTATIONID` channel variable
	RemoteStationID string
	// PagesTransferred The number of pages that have been transferred
	PagesTransferred string
	// Resolution The negotiated resolution
	Resolution string
	// TransferRate The negotiated transfer rate
	TransferRate string
	// FileName The files being affected by the fax operation
	FileName string
}

func (SendFAXEvent) EventTypeName() string {
	return "SendFAX"
}
