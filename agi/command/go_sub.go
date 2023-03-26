package command

type GoSubCommand struct {
	Context          string
	Extension        string
	Priority         int
	OptionalArgument *string
}

func (cmd GoSubCommand) CommandString() string {
	return joinCommand([]interface{}{"GOSUB", cmd.Context, cmd.Extension, cmd.Priority, cmd.OptionalArgument})
}

func (cmd GoSubCommand) SetOptionalArgument(arg string) Command {
	cmd.OptionalArgument = &arg
	return cmd
}

func NewGoSubCommand(context, extension string, priority int) Command {
	return GoSubCommand{
		Context:   context,
		Extension: extension,
		Priority:  priority,
	}
}
