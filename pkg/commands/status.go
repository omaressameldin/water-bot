package commands

import (
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
)

func AddStatusCmd(bot *slacker.Slacker) {
	bot.Command(StatusCmd, statusCmd())
}

func statusCmd() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description: StatusDescription,
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			channel := request.Event().Channel
			rtm := response.RTM()

			statusQuestion(rtm, channel)
		},
	}
}

func statusQuestion(rtm *slack.RTM, channel string) {
	attachment := slack.Attachment{
		Text:       STATUS_QUESTION,
		Color:      "#f9a41b",
		CallbackID: STATUS_CALLBACK_ID,
		Actions: []slack.AttachmentAction{
			{
				Name:  STATUS_CONFIRM_VAL,
				Text:  STATUS_CONFIRM_TEXT,
				Type:  "button",
				Style: "primary",
			},
			{
				Name:  STATUS_CANCEL_VAL,
				Text:  STATUS_CANCEL_TEXT,
				Type:  "button",
				Style: "danger",
			},
		},
	}
	rtm.PostMessage(channel, slack.MsgOptionAttachments(attachment))
}
