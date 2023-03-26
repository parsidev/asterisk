package command

type SetContextCommand struct {
	DesiredContext string
}

func (cmd SetContextCommand) CommandString() string {
	return joinCommand([]interface{}{"SET CONTEXT", cmd.DesiredContext})
}

func NewSetContextCommand(context string) Command {
	return SetContextCommand{DesiredContext: context}
}
