package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/water-bot/internal/attachments"
	"github.com/omaressameldin/water-bot/internal/env"
	"github.com/omaressameldin/water-bot/internal/utils"
	"github.com/shomali11/slacker"
)

type Reply struct {
	Attachments []slack.Attachment `json:"attachments"`
	Text        string             `json:"text"`
}

func HandleActions(bot *slacker.Slacker) {
	port, err := env.GetActionPort()
	if err != nil {
		log.Fatalf("need to add %s to .env file", env.ACTIONS_PORT_KEY)
	}
	http.HandleFunc(ROUTE, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			bot.Client().ConnectRTM()
			payload, err := unmarshalPayload(r)
			utils.HttpError(err, "error responsing to action", w)
			switch payload.CallbackID {
			case attachments.ORDER_START_CALLBACK_ID:
				{
					firstChoice(*payload, w)
				}
			case attachments.ORDER_STAGE_1_CALLBACK_ID:
				{
					secondChoice(*payload, w)
				}
			case attachments.ORDER_STAGE_2_CALLBACK_ID:
				{
					confirmOrder(*payload, w)
				}
			}

		}
	})
	log.Printf("listening on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func unmarshalPayload(r *http.Request) (*slack.InteractionCallback, error) {
	var payload slack.InteractionCallback
	err := json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func sendReply(w http.ResponseWriter, r Reply) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func mightCancel(
	payload slack.InteractionCallback,
	w http.ResponseWriter,
	text string,
) bool {

	isCancel := payload.ActionCallback.AttachmentActions[0].Name == attachments.CANCEL_VAL
	if isCancel {
		sendReply(w, Reply{
			Attachments: []slack.Attachment{slack.Attachment{
				Text:  text,
				Color: "#dc3545",
			}},
		})
	}

	return isCancel
}
