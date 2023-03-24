package actions

import "github.com/parsidev/asterisk/ami/message"

type QueueMemberRingInUseAction struct {
	ActionID  string
	Interface string
	RingInUse string
	Queue     string
}

func (QueueMemberRingInUseAction) ActionTypeName() string {
	return "QueueMemberRingInUse"
}
func (a QueueMemberRingInUseAction) GetActionID() string {
	return a.ActionID
}
func (a *QueueMemberRingInUseAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) QueueMemberRingInUse(iface string, ringInUse string, opts ...message.RequestOption) (res *message.Response, err error) {
	req := &QueueMemberRingInUseAction{
		Interface: iface,
		RingInUse: ringInUse,
	}
	res = &message.Response{}
	return res, cli.Action(req, res, opts...)
}
