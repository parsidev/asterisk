package command

type SetCallerIdCommand struct {
	Number string
}

func (cmd SetCallerIdCommand) CommandString() string {
	return joinCommand([]interface{}{"SET CALLERID", cmd.Number})
}

func NewSetCallerIdCommand(number string) Command {
	return SetCallerIdCommand{Number: number}
}
