package command

type ExecCommand struct {
	Application string
	Options     string
}

func (cmd ExecCommand) CommandString() string {
	return joinCommand([]interface{}{"EXEC", cmd.Application, cmd.Options})
}

func NewExecCommand(application, options string) Command {
	return ExecCommand{Application: application, Options: options}
}
