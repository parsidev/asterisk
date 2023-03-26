package command

type GetDataCommand struct {
	File      string
	Timeout   *int
	MaxDigits *string
}

func (cmd GetDataCommand) CommandString() string {
	return joinCommand([]interface{}{"GET DATA", cmd.File, cmd.Timeout, cmd.MaxDigits})
}

func (cmd GetDataCommand) SetTimeout(v int) Command {
	cmd.Timeout = &v
	return cmd
}
func (cmd GetDataCommand) SetMaxDigits(v string) Command {
	cmd.MaxDigits = &v
	return cmd
}

func NewGetDataCommand(file string) Command {
	return GetDataCommand{File: file}
}
