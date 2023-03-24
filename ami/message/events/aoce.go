package events

type AOCEEvent struct {
	ChargingAssociation string
	Number              string
	Plan                string
	Id                  string
	Charge              string
	Type                string
	BillingID           string
	TotalType           string
	Currency            string
	Name                string
	Cost                string
	Multiplier          string
	Units               string
	NumberOf            string
	TypeOf              string
}

func (AOCEEvent) EventTypeName() string {
	return "AOC-E"
}
