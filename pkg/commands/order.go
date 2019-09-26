package commands

import (
	"github.com/nlopes/slack"
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
		CallbackID: ORDER_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			{
				Name:  ORDER_START_VAL,
				Text:  ORDER_START_TEXT,
				Type:  "button",
				Style: "primary",
			},
			{
				Name:  ORDER_CANCEL_VAL,
				Text:  ORDER_CANCEL_TEXT,
				Type:  "button",
				Style: "danger",
			},
		},
	}
	rtm.PostMessage(channel, slack.MsgOptionAttachments(attachment))
}
