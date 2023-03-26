package command

type SetMusicCommand struct {
	Class string
}

func (cmd SetMusicCommand) CommandString() string {
	return joinCommand([]interface{}{"SET MUSIC", cmd.Class})
}

func NewSetMusicCommand(class string) Command {
	return SetMusicCommand{Class: class}
}
