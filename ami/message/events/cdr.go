package events

type CdrEvent struct {
	AccountCode        string
	Source             string
	Destination        string
	DestinationContext string
	CallerID           string
	Channel            string
	DestinationChannel string
	LastApplication    string
	LastData           string
	StartTime          string
	AnswerTime         string
	EndTime            string
	Duration           string
	BillableSeconds    string
	Disposition        string
	AMAFlags           string
	UniqueID           string
	UserField          string
}

func (CdrEvent) EventTypeName() string {
	return "Cdr"
}
