package command

type HangupCommand struct {
	ChannelName *string
}

func (cmd HangupCommand) CommandString() string {
	return joinCommand([]interface{}{"HANGUP", cmd.ChannelName})
}

func (cmd HangupCommand) SetChannelName(v string) Command {
	cmd.ChannelName = &v
	return cmd
}

func NewHangupCommand() Command {
	return HangupCommand{}
}
