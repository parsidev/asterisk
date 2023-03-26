package command

type DatabaseDelTreeCommand struct {
	Family  string
	KeyTree *string
}

func (cmd DatabaseDelTreeCommand) CommandString() string {
	return joinCommand([]interface{}{"DATABASE DELTREE", cmd.Family, cmd.KeyTree})
}

func (cmd DatabaseDelTreeCommand) SetKeyTree(v string) Command {
	cmd.KeyTree = &v
	return cmd
}

func NewDatabaseDelTreeCommand(family string) Command {
	return DatabaseDelTreeCommand{Family: family}
}
