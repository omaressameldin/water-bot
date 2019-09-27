package actions

import (
	"net/http"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/water-bot/internal/attachments"
	"github.com/omaressameldin/water-bot/internal/utils"
)

func firstChoice(payload slack.InteractionCallback, w http.ResponseWriter) {
	mightCancel(payload, w, CANCEL_TEXT)

	addChoice(
		w,
		CHOICE_QUESTION_1,
		attachments.ORDER_STAGE_1_CALLBACK_ID,
		CHOICE_TEXT_1,
		CHOICE_VAL_1,
	)
}

func secondChoice(payload slack.InteractionCallback, w http.ResponseWriter) {
	mightCancel(payload, w, CANCEL_TEXT)

	addChoice(
		w,
		CHOICE_QUESTION_2,
		attachments.ORDER_STAGE_2_CALLBACK_ID,
		CHOICE_TEXT_2,
		CHOICE_VAL_2,
	)
}

func confirmOrder(payload slack.InteractionCallback, w http.ResponseWriter) {
	mightCancel(payload, w, CANCEL_TEXT)

	sendReply(w, Reply{
		Attachments: []slack.Attachment{slack.Attachment{
			Text:  CONFIRM_TEXT,
			Color: "#28a745",
		}},
	})
}

func addChoice(
	w http.ResponseWriter,
	question string,
	callbackId string,
	selectText string,
	selectVal string,
) {
	waterOptions, err := attachments.GenerateNumberOptions(0, 10)
	if err != nil {
		utils.HttpError(err, "Cant order water", w)
	}

	waterOrder := slack.Attachment{
		Text:       question,
		Color:      "#17a2b8",
		CallbackID: callbackId,
		Actions: []slack.AttachmentAction{
			attachments.Select(
				waterOptions,
				selectText,
				selectVal,
			),
			attachments.CancelButton("Cancel Order"),
		},
	}
	sendReply(w, Reply{
		Attachments: []slack.Attachment{waterOrder},
	})
}
