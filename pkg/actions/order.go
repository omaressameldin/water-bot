package actions

import (
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/omaressameldin/water-bot/pkg/commands"
)

func orderCallback(payload slack.InteractionCallback, w http.ResponseWriter) {
	switch payload.ActionCallback.AttachmentActions[0].Name {
	case commands.ORDER_CONFIRM_VAL:
		{
			confirmAction(w)
		}
	default:
		cancelAction(w)
	}

}

func confirmAction(w http.ResponseWriter) {
	w.Write([]byte("confirmed"))
}

func cancelAction(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Reply{
		Attachments: []slack.Attachment{slack.Attachment{
			Text:  CANCEL_TEXT,
			Color: "#dc3545",
		}},
	})
}
