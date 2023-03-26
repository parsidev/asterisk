package command

type VerboseCommand struct {
	Message string
	Level   string
}

func (cmd VerboseCommand) CommandString() string {
	return joinCommand([]interface{}{"VERBOSE", cmd.Message, cmd.Level})
}

func NewVerboseCommand(message, level string) Command {
	return VerboseCommand{Message: message, Level: level}
}
