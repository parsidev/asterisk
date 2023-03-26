package command

type SayPhoneticCommand struct {
	String       string
	EscapeDigits string
}

func (cmd SayPhoneticCommand) CommandString() string {
	return joinCommand([]interface{}{"SAY PHONETIC", cmd.String, cmd.EscapeDigits})
}

func NewSayPhoneticCommand(string, escapeDigits string) Command {
	return SayPhoneticCommand{String: string, EscapeDigits: escapeDigits}
}
