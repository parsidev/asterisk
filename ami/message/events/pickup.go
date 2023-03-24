package events

type PickupEvent struct {
}

func (PickupEvent) EventTypeName() string {
	return "Pickup"
}
