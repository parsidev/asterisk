package command

type ChannelStatusCommand struct {
	ChannelName *string
}

func (cmd ChannelStatusCommand) CommandString() string {
	return joinCommand([]interface{}{"CHANNEL STATUS", cmd.ChannelName})
}

func (cmd ChannelStatusCommand) SetChannelName(v string) Command {
	cmd.ChannelName = &v
	return cmd
}

func NewChannelStatusCommand() Command {
	return ChannelStatusCommand{}
}
