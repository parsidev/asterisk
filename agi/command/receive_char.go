package command

type ReceiveCharCommand struct {
	Timeout int
}

func (cmd ReceiveCharCommand) CommandString() string {
	return joinCommand([]interface{}{"RECEIVE CHAR", cmd.Timeout})
}

func NewReceiveCharCommand(timeout int) Command {
	return ReceiveCharCommand{Timeout: timeout}
}
