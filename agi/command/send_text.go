package command

type SendTextCommand struct {
	TextToSend string
}

func (cmd SendTextCommand) CommandString() string {
	return joinCommand([]interface{}{"SEND TEXT", cmd.TextToSend})
}

func NewSendTextCommand(text string) Command {
	return SendTextCommand{TextToSend: text}
}
