package events

type StatusEvent struct {
	ActionID                   string
	Type                       string
	DNID                       string
	EffectiveConnectedLineNum  string
	EffectiveConnectedLineName string
	TimeToHangup               string
	BridgeID                   string
	Application                string
	Data                       string
	Nativeformats              string
	Readformat                 string
	Readtrans                  string
	Writeformat                string
	Writetrans                 string
	Callgroup                  string
	Pickupgroup                string
	Seconds                    string
}

func (StatusEvent) EventTypeName() string {
	return "Status"
}
