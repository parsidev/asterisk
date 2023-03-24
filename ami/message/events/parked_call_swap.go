package events

type ParkedCallSwapEvent struct {
	ParkerDialString string
	Parkinglot       string
	ParkingSpace     string
	ParkingTimeout   string
	ParkingDuration  string
}

func (ParkedCallSwapEvent) EventTypeName() string {
	return "ParkedCallSwap"
}
