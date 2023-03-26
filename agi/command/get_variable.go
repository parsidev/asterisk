package command

type GetVariableCommand struct {
	VariableName string
}

func (cmd GetVariableCommand) CommandString() string {
	return joinCommand([]interface{}{"GET VARIABLE", cmd.VariableName})
}

func NewGetVariableCommand(variable string) Command {
	return GetVariableCommand{VariableName: variable}
}
