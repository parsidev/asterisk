package events

type CELEvent struct {
	EventName     string
	AccountCode   string
	CallerIDnum   string
	CallerIDName  string
	CallerIDani   string
	CallerIDrdnis string
	CallerIDdnid  string
	Exten         string
	Context       string
	Application   string
	AppData       string
	EventTime     string
	AMAFlags      string
	UniqueID      string
	LinkedID      string
	UserField     string
	Peer          string
	PeerAccount   string
	Extra         string
}

func (CELEvent) EventTypeName() string {
	return "CEL"
}
