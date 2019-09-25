package commands

import (
	"log"
	"time"

	"github.com/omaressameldin/water-bot/internal/utils"
	"github.com/shomali11/slacker"
)

func AddInitCmd(bot *slacker.Slacker) {
	bot.Init(checkwaterDaily(bot))
}

func checkwaterDaily(bot *slacker.Slacker) func() {
	return func() {
		bot.Client().ConnectRTM()
		for {
			d, e := utils.TimeTill(utils.NextTime(INIT_HOUR, INIT_MIN, INIT_SEC))
			if e != nil {
				log.Fatal(e)
			}

			time.Sleep(*d)

			orderQuestion(bot.RTM(), INIT_POST_CHANNEL)
		}
	}
}
