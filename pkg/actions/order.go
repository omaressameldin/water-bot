package actions

import (
	"log"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/water-bot/internal/utils"
	"github.com/omaressameldin/water-bot/pkg/commands"
)

func orderStartCallback(payload slack.InteractionCallback, w http.ResponseWriter) {
	switch payload.ActionCallback.AttachmentActions[0].Name {
	case commands.ORDER_START_VAL:
		{
			confirmChoice(
				w,
				CHOICE_QUESTION_1,
				CHOICE_CALLBACK_ID_1,
				CHOICE_TEXT_1,
				CHOICE_VAL_1,
			)
		}
	default:
		cancelAction(w)
	}
}

func firstChoiceCallback(payload slack.InteractionCallback, w http.ResponseWriter) {
	confirmChoice(
		w,
		CHOICE_QUESTION_2,
		CHOICE_CALLBACK_ID_2,
		CHOICE_TEXT_2,
		CHOICE_VAL_2,
	)
}

func secondChoiceCallback(payload slack.InteractionCallback, w http.ResponseWriter) {
	log.Println(len(payload.ActionCallback.AttachmentActions))
	sendReply(w, Reply{
		Attachments: []slack.Attachment{slack.Attachment{
			Text:  CONFIRM_TEXT,
			Color: "#28a745",
		}},
	})
}

func confirmChoice(
	w http.ResponseWriter,
	question string,
	callbackId string,
	selectText string,
	selectVal string,
) {
	waterOptions, err := utils.GenerateAttachmentOptions(0, 10)
	if err != nil {
		utils.HttpError(err, "Cant order water", w)
	}

	waterOrder := slack.Attachment{
		Text:       question,
		Color:      "#17a2b8",
		CallbackID: callbackId,
		Actions: []slack.AttachmentAction{
			{
				Name:    selectVal,
				Text:    selectText,
				Type:    "select",
				Options: waterOptions,
			},
		},
	}
	sendReply(w, Reply{
		Attachments: []slack.Attachment{waterOrder},
	})
}

func cancelAction(w http.ResponseWriter) {
	sendReply(w, Reply{
		Attachments: []slack.Attachment{slack.Attachment{
			Text:  CANCEL_TEXT,
			Color: "#dc3545",
		}},
	})
}
