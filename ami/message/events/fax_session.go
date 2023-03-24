package events

type FAXSessionEvent struct {
	ActionID            string
	SessionNumber       string
	Operation           string
	State               string
	ErrorCorrectionMode string
	DataRate            string
	ImageResolution     string
	PageNumber          string
	FileName            string
	PagesTransmitted    string
	PagesReceived       string
	TotalBadLines       string
}

func (FAXSessionEvent) EventTypeName() string {
	return "FAXSession"
}
