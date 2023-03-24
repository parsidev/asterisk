package events

type UnParkedCallEvent struct {
	// ParkerDialString Dial String that can be used to call back the parker on ParkingTimeout.
	ParkerDialString string
	// Parkinglot Name of the parking lot that the parkee is parked in
	Parkinglot string
	// ParkingSpace Parking Space that the parkee is parked in
	ParkingSpace string
	// ParkingTimeout Time remaining until the parkee is forcefully removed from parking in seconds
	ParkingTimeout string
	// ParkingDuration Time the parkee has been in the parking bridge (in seconds)
	ParkingDuration string
}

func (UnParkedCallEvent) EventTypeName() string {
	return "UnParkedCall"
}
