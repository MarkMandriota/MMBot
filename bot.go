package main

import (
	"errors"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

type Bot struct {
	session *dg.Session
	*Config

	errors chan error
}

func (b *Bot) OpenErrHandler() error {
	if b.errors == nil {
		return errors.New("error channel not initialized")
	}

	for {
		switch err := <-b.errors; err.(type) {
		case errors.NotEnoughPermissions:

		}
	}
}

func (b *Bot) MessageCreate(s *dg.Session, m *dg.MessageCreate) {
	msgLogs.WriteString(fmt.Sprintf("%v By %s sent %s\n", m.Timestamp, m.Author.String(), m.Content))

	if !strings.HasPrefix(m.Content, bot.Prefix) || m.Author.Bot {
		return
	}

	args := strings.Fields(strings.TrimPrefix(m.Content, bot.Prefix))
	if len(args) < 1 {
		return
	}

	for _, p := range cmd[args[0]].Perm() {
		userP, _ := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

		if userP&p == p {
			cmd[args[0]].Exec(m)
		}
	}
}
