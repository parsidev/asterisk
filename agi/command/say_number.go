package command

type SayNumberCommand struct {
	Number       string
	EscapeDigits string
	Gender       *string
}

func (cmd SayNumberCommand) CommandString() string {
	return joinCommand([]interface{}{"SAY NUMBER", cmd.Number, cmd.EscapeDigits, cmd.Gender})
}

func (cmd SayNumberCommand) SetGender(v string) Command {
	cmd.Gender = &v
	return cmd
}

func NewSayNumberCommand(number, escapeDigits string) Command {
	return SayNumberCommand{Number: number, EscapeDigits: escapeDigits}
}
