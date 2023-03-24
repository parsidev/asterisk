package events

type ParkedCallEvent struct {
	ParkerDialString string
	Parkinglot       string
	ParkingSpace     string
	ParkingTimeout   string
	ParkingDuration  string
}

func (ParkedCallEvent) EventTypeName() string {
	return "ParkedCall"
}
