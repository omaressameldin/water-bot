package actions

import (
	"net/http"
	"strconv"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/water-bot/internal/attachments"
	"github.com/omaressameldin/water-bot/internal/utils"
	"github.com/omaressameldin/water-bot/pkg/automation"
)

func firstChoice(payload slack.InteractionCallback, w http.ResponseWriter) {
	if mightCancel(payload, w, CANCEL_TEXT) {
		return
	}

	answer, err := attachments.MarshalSelectedOptions([]attachments.SelectedOption{
		{
			Name: CHOICE_VAL_1,
		},
	})
	utils.HttpError(err, "Error sending reply", w)

	addChoice(
		w,
		CHOICE_QUESTION_1,
		attachments.ORDER_STAGE_1_CALLBACK_ID,
		CHOICE_TEXT_1,
		answer,
	)
}

func secondChoice(payload slack.InteractionCallback, w http.ResponseWriter) {
	if mightCancel(payload, w, CANCEL_TEXT) {
		return
	}

	answers, err := attachments.AddAnswer(
		payload.ActionCallback.AttachmentActions[0].Name,
		payload.ActionCallback.AttachmentActions[0].SelectedOptions[0].Value,
	)
	utils.HttpError(err, "Error sending reply", w)

	newAnswers := append(answers, attachments.SelectedOption{
		Name: CHOICE_VAL_2,
	})

	val, err := attachments.MarshalSelectedOptions(newAnswers)
	utils.HttpError(err, "something went wrong please try again", w)

	addChoice(
		w,
		CHOICE_QUESTION_2,
		attachments.ORDER_STAGE_2_CALLBACK_ID,
		CHOICE_TEXT_2,
		val,
	)
}

func confirmOrder(payload slack.InteractionCallback, w http.ResponseWriter) {
	if mightCancel(payload, w, CANCEL_TEXT) {
		return
	}

	answers, err := attachments.AddAnswer(
		payload.ActionCallback.AttachmentActions[0].Name,
		payload.ActionCallback.AttachmentActions[0].SelectedOptions[0].Value,
	)
	utils.HttpError(err, "Error sending reply", w)

	stillWaterBoxes, err := strconv.Atoi(answers[0].Val)
	utils.HttpError(err, "Error sending reply", w)

	sparklingWaterBoxes, err := strconv.Atoi(answers[1].Val)
	utils.HttpError(err, "Error sending reply", w)
	automation.OrderWater(stillWaterBoxes, sparklingWaterBoxes)

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
