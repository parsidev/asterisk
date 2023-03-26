package command

type DatabaseGetCommand struct {
	Family string
	Key    string
}

func (cmd DatabaseGetCommand) CommandString() string {
	return joinCommand([]interface{}{"DATABASE GET", cmd.Family, cmd.Key})
}

func NewDatabaseGetCommand(family, Key string) Command {
	return DatabaseGetCommand{Family: family, Key: Key}
}
