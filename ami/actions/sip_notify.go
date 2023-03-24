package actions

type SIPnotifyAction struct {
	ActionID string
	Channel  string
	Variable string
	CallID   string
}

func (SIPnotifyAction) ActionTypeName() string {
	return "SIPnotify"
}
func (a SIPnotifyAction) GetActionID() string {
	return a.ActionID
}
func (a *SIPnotifyAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) SIPnotify(channel string, variable string, opts ...RequestOption) (res *Response, err error) {
	req := &SIPnotifyAction{
		Channel:  channel,
		Variable: variable,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
