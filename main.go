package main

import (
	"log"

	"github.com/omaressameldin/water-bot/pkg/actions"
	"github.com/omaressameldin/water-bot/pkg/bot"
)

func main() {
	b, err := bot.CreateBot()
	if err != nil {
		log.Fatal(err)
	}
	go (func() {
		actions.HandleActions(b.SlackBot)
	})()
	b.StartListening()
}
