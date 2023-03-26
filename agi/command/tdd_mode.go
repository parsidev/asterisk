package command

type TddModeCommand struct {
	Boolean string
}

func (cmd TddModeCommand) CommandString() string {
	return joinCommand([]interface{}{"TDD MODE", cmd.Boolean})
}

func NewTddModeCommand(boolean string) Command {
	return TddModeCommand{Boolean: boolean}
}
