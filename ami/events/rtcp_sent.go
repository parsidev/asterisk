package events

type RTCPSentEvent struct {
	SSRC                        string
	PT                          string
	To                          string
	ReportCount                 string
	SentNTP                     string
	SentRTP                     string
	SentPackets                 string
	SentOctets                  string
	ReportXSourceSSRC           string
	ReportXFractionLost         string
	ReportXCumulativeLost       string
	ReportXHighestSequence      string
	ReportXSequenceNumberCycles string
	ReportXIAJitter             string
	ReportXLSR                  string
	ReportXDLSR                 string
}

func (RTCPSentEvent) EventTypeName() string {
	return "RTCPSent"
}
