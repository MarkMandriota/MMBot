package cmds

import dg "github.com/bwmarrin/discordgo"

type Command interface {
	Name() string
	Info(a ...string) string
	Exec(*dg.Session, *dg.MessageCreate) error
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"ping": &Ping{},
	}
}
