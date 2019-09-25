package main

import (
	"log"

	"github.com/omaressameldin/water-bot/pkg/bot"
)

func main() {
	b, err := bot.CreateBot()
	if err != nil {
		log.Fatal(err)
	}

	b.StartListening()

}
