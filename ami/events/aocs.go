package events

type AOCSEvent struct {
	Chargeable   string
	RateType     string
	Currency     string
	Name         string
	Cost         string
	Multiplier   string
	ChargingType string
	StepFunction string
	Granularity  string
	Length       string
	Scale        string
	Unit         string
	SpecialCode  string
}

func (AOCSEvent) EventTypeName() string {
	return "AOC-S"
}
