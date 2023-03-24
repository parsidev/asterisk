package events

type DeviceStateChangeEvent struct {
	Device string
	State  string
}

func (DeviceStateChangeEvent) EventTypeName() string {
	return "DeviceStateChange"
}
