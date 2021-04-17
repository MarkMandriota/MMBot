package main

import (
	"encoding/json"
	"log"
	"os"

	dg "github.com/bwmarrin/discordgo"
)

const (
	_CONFIG_PASS = "config.json"
	_LOGS_PASS   = "logs/"
)

var (
	err error

	bot     *Bot
	msgLogs *os.File
)

func init() {
	bot = &Bot{}

	msgLogs, err = os.OpenFile(_LOGS_PASS+"msg.logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
}

func init() {
	file, err := os.Open(_CONFIG_PASS)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(bot); err != nil {
		panic(err)
	}
}

func main() {
	defer msgLogs.Close()

	if bot.session, err = dg.New(bot.Token); err != nil {
		panic(err)
	}

	bot.session.AddHandler(bot.MessageCreate)

	if err = bot.session.Open(); err != nil {
		panic(err)
	}
	defer bot.session.Close()

	log.Println("Bot's running...")
	<-make(chan struct{})
}
