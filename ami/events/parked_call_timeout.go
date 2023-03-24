package events

type ParkedCallTimeOutEvent struct {
	ParkerDialString string
	Parkinglot       string
	ParkingSpace     string
	ParkingTimeout   string
	ParkingDuration  string
}

func (ParkedCallTimeOutEvent) EventTypeName() string {
	return "ParkedCallTimeOut"
}
