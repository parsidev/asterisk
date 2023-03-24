package events

type FAXStatusEvent struct {
	Operation string
	// Status A text message describing the current status of the fax
	Status string
	// LocalStationID The value of the `LOCALSTATIONID` channel variable
	LocalStationID string
	// FileName The files being affected by the fax operation
	FileName string
}

func (FAXStatusEvent) EventTypeName() string {
	return "FAXStatus"
}
