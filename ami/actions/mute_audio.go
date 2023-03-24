package actions

type MuteAudioAction struct {
	ActionID  string
	Channel   string
	Direction string
	State     string
}

func (MuteAudioAction) ActionTypeName() string {
	return "MuteAudio"
}
func (a MuteAudioAction) GetActionID() string {
	return a.ActionID
}
func (a *MuteAudioAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) MuteAudio(channel string, direction string, state string, opts ...RequestOption) (res *Response, err error) {
	req := &MuteAudioAction{
		Channel:   channel,
		Direction: direction,
		State:     state,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
