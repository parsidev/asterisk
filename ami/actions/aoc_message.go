package actions

type AOCMessageAction struct {
	ActionID                  string
	Channel                   string
	ChannelPrefix             string
	MsgType                   string
	ChargeType                string
	UnitAmount0               string
	UnitType0                 string
	CurrencyName              string
	CurrencyAmount            string
	CurrencyMultiplier        string
	TotalType                 string
	AOCBillingId              string
	ChargingAssociationId     string
	ChargingAssociationNumber string
	ChargingAssociationPlan   string
}

func (AOCMessageAction) ActionTypeName() string {
	return "AOCMessage"
}
func (a AOCMessageAction) GetActionID() string {
	return a.ActionID
}
func (a *AOCMessageAction) SetActionID(actionID string) {
	a.ActionID = actionID
}
func (cli *Client) AOCMessage(channel string, msgType string, chargeType string, opts ...RequestOption) (res *Response, err error) {
	req := &AOCMessageAction{
		Channel:    channel,
		MsgType:    msgType,
		ChargeType: chargeType,
	}
	res = &Response{}
	return res, cli.Action(req, res, opts...)
}
