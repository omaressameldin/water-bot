package commands

import (
	"github.com/nlopes/slack"
	"github.com/omaressameldin/water-bot/internal/attachments"
	"github.com/shomali11/slacker"
)

func AddOrderCmd(bot *slacker.Slacker) {
	bot.Command(ORDER_CMD, orderCmd())
}

func orderCmd() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: ORDER_DESCRIPTION,
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			channel := request.Event().Channel
			rtm := response.RTM()

			orderQuestion(rtm, channel)
		},
	}
}

func orderQuestion(rtm *slack.RTM, channel string) {
	attachment := slack.Attachment{
		Text:       ORDER_QUESTION,
		Color:      "#f9a41b",
		CallbackID: attachments.ORDER_START_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			attachments.ConfirmButton(ORDER_START_TEXT),
			attachments.CancelButton(ORDER_CANCEL_TEXT),
		},
	}
	rtm.PostMessage(channel, slack.MsgOptionAttachments(attachment))
}
