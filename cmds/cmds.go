package cmds

import dg "github.com/bwmarrin/discordgo"

type Command interface {
	Name() string
	Info(a ...string) string
	Perm() []int
	Exec(s *dg.Session, m *dg.MessageCreate)
}

func NewCommands(s *dg.Session) map[string]Command {
	return map[string]Command{
		"ping": &Ping{s},
	}
}
