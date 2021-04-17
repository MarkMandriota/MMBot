package main

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

type Bot struct {
	session *dg.Session
	*Config
}

func (b *Bot) MessageCreate(s *dg.Session, m *dg.MessageCreate) {
	msgLogs.WriteString(fmt.Sprintf("%v By %s sent %s\n", m.Timestamp, m.Author.String(), m.Content))
}
