package events

type MCIDEvent struct {
	MCallerIDNumValid       string
	MCallerIDNum            string
	MCallerIDton            string
	MCallerIDNumPlan        string
	MCallerIDNumPres        string
	MCallerIDNameValid      string
	MCallerIDName           string
	MCallerIDNameCharSet    string
	MCallerIDNamePres       string
	MCallerIDSubaddr        string
	MCallerIDSubaddrType    string
	MCallerIDSubaddrOdd     string
	MCallerIDPres           string
	MConnectedIDNumValid    string
	MConnectedIDNum         string
	MConnectedIDton         string
	MConnectedIDNumPlan     string
	MConnectedIDNumPres     string
	MConnectedIDNameValid   string
	MConnectedIDName        string
	MConnectedIDNameCharSet string
	MConnectedIDNamePres    string
	MConnectedIDSubaddr     string
	MConnectedIDSubaddrType string
	MConnectedIDSubaddrOdd  string
	MConnectedIDPres        string
}

func (MCIDEvent) EventTypeName() string {
	return "MCID"
}
