package command

type GetOptionCommand struct {
	FileName     string
	EscapeDigits string
	Timeout      *int
}

func (cmd GetOptionCommand) CommandString() string {
	return joinCommand([]interface{}{"GET OPTION", cmd.FileName, cmd.EscapeDigits, cmd.Timeout})
}

func (cmd GetOptionCommand) SetTimeout(v int) Command {
	cmd.Timeout = &v
	return cmd
}

func NewGetOptionCommand(fileName string) Command {
	return GetOptionCommand{FileName: fileName}
}
