package command

type ReceiveTextCommand struct {
	Timeout int
}

func (cmd ReceiveTextCommand) CommandString() string {
	return joinCommand([]interface{}{"RECEIVE TEXT", cmd.Timeout})
}

func NewReceiveTextCommand(timeout int) Command {
	return ReceiveTextCommand{Timeout: timeout}
}
