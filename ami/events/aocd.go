package events

type AOCDEvent struct {
	Charge     string
	Type       string
	BillingID  string
	TotalType  string
	Currency   string
	Name       string
	Cost       string
	Multiplier string
	Units      string
	NumberOf   string
	TypeOf     string
}

func (AOCDEvent) EventTypeName() string {
	return "AOC-D"
}
