package actions

type OriginateAction struct {
	ActionID       string
	Channel        string
	Exten          string
	Context        string
	Priority       int
	Application    string
	Data           string
	Timeout        int
	CallerID       string
	Variable       string
	Account        string
	EarlyMedia     string
	Async          string
	Codecs         string
	ChannelId      string
	OtherChannelId string
}

func (OriginateAction) ActionTypeName() string {
	return "Originate"
}
func (a OriginateAction) GetActionID() string {
	return a.ActionID
}
func (a *OriginateAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) Originate(channel string, opts ...RequestOption) (res *Response, err error) {
	req := &OriginateAction{
		Channel: channel,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
