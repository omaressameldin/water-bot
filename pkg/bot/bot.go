package bot

import (
	"context"
	"log"

	"github.com/omaressameldin/water-bot/internal/env"
	"github.com/omaressameldin/water-bot/pkg/commands"
	"github.com/shomali11/slacker"
)

type Bot struct {
	SlackBot *slacker.Slacker
}

func CreateBot() (*Bot, error) {
	t, err := env.GetToken()
	if err != nil {
		return nil, err
	}
	return &Bot{
		SlackBot: slacker.NewClient(t),
	}, nil
}

func (b *Bot) StartListening() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	commands.AddInitCmd(b.SlackBot)
	commands.AddOrderCmd(b.SlackBot)

	log.Println("Listening...")
	err := b.SlackBot.Listen(ctx)
	if err != nil {
		return err
	}

	return nil
}
