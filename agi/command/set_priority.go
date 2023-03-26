package command

type SetPriorityCommand struct {
	Priority int
}

func (cmd SetPriorityCommand) CommandString() string {
	return joinCommand([]interface{}{"SET PRIORITY", cmd.Priority})
}

func NewSetPriorityCommand(priority int) Command {
	return SetPriorityCommand{Priority: priority}
}
