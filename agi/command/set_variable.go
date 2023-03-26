package command

type SetVariableCommand struct {
	VariableName string
	Value        string
}

func (cmd SetVariableCommand) CommandString() string {
	return joinCommand([]interface{}{"SET VARIABLE", cmd.VariableName, cmd.Value})
}

func NewSetVariableCommand(variable, value string) Command {
	return SetVariableCommand{VariableName: variable, Value: value}
}
