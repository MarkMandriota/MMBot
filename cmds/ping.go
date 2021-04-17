package cmds

import (
	"fmt"
	"time"

	dg "github.com/bwmarrin/discordgo"
)

type Ping struct {
	s *dg.Session
}

func (Ping) Name() string {
	return "ping"
}

func (Ping) Info(...string) string {
	return "return bot ping"
}

func (Ping) Perm() []int {
	return []int{0}
}

func (p *Ping) Exec(m *dg.MessageCreate) {
	start, _ := m.Timestamp.Parse()
	duration := time.Since(start)

	p.s.ChannelMessageSendEmbed(m.ChannelID, &dg.MessageEmbed{
		Author: &dg.MessageEmbedAuthor{
			URL:  m.Author.Avatar,
			Name: m.Author.String(),
		},
		Title:       "Pong!",
		Description: fmt.Sprint(duration),
	})
}
