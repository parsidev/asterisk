package command

type WaitForDigitCommand struct {
	Timeout int
}

func (cmd WaitForDigitCommand) CommandString() string {
	return joinCommand([]interface{}{"WAIT FOR DIGIT", cmd.Timeout})
}

func NewWaitForDigitCommand(timeout int) Command {
	return WaitForDigitCommand{Timeout: timeout}
}
