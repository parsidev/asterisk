package command

type SendImageCommand struct {
	Image string
}

func (cmd SendImageCommand) CommandString() string {
	return joinCommand([]interface{}{"SEND IMAGE", cmd.Image})
}

func NewSendImageCommand(image string) Command {
	return SendImageCommand{Image: image}
}
