package events

type RTCPReceivedEvent struct {
	SSRC                        string
	PT                          string
	From                        string
	RTT                         string
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

func (RTCPReceivedEvent) EventTypeName() string {
	return "RTCPReceived"
}
