package command

type DatabasePutCommand struct {
	Family string
	Key    string
	Value  string
}

func (cmd DatabasePutCommand) CommandString() string {
	return joinCommand([]interface{}{"DATABASE PUT", cmd.Family, cmd.Key, cmd.Value})
}

func NewDatabasePutCommand(family, key, value string) Command {
	return DatabasePutCommand{Family: family, Key: key, Value: value}
}
