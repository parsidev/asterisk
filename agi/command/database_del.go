package command

type DatabaseDelCommand struct {
	Family string
	Key    string
}

func (cmd DatabaseDelCommand) CommandString() string {
	s := []interface{}{"DATABASE DEL", cmd.Family, cmd.Key}
	return joinCommand(s)
}

func NewDatabaseDelCommand(family, key string) Command {
	return DatabaseDelCommand{Family: family, Key: key}
}
