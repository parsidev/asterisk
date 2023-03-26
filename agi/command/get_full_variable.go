package command

type GetFullVariableCommand struct {
	Expression  string
	ChannelName *string
}

func (cmd GetFullVariableCommand) CommandString() string {
	return joinCommand([]interface{}{"GET FULL VARIABLE", cmd.Expression, cmd.ChannelName})
}

func (cmd GetFullVariableCommand) SetChannelName(v string) Command {
	cmd.ChannelName = &v
	return cmd
}

func NewGetFullVariableCommand(expression string) Command {
	return GetFullVariableCommand{Expression: expression}
}
