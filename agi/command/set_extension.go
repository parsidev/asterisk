package command

type SetExtensionCommand struct {
	NewExtension string
}

func (cmd SetExtensionCommand) CommandString() string {
	return joinCommand([]interface{}{"SET EXTENSION", cmd.NewExtension})
}

func NewSetExtensionCommand(newExtension string) Command {
	return SetExtensionCommand{NewExtension: newExtension}
}
