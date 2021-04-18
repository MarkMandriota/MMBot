package cmds

import (
	dg "github.com/bwmarrin/discordgo"
)

type Ping struct{}

func (Ping) Name() string {
	return "ping"
}

func (Ping) Info(...string) string {
	return "returns bot ping"
}

func (p *Ping) Exec(s *dg.Session, m *dg.MessageCreate) (err error) {
	//start, _ := m.Timestamp.Parse()
	//duration := time.Since(start)

	s.ChannelMessageSend(m.ChannelID, "Pong!")
	return
}
