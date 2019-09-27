package attachments

import (
	"fmt"

	"github.com/nlopes/slack"
)

func GenerateNumberOptions(begin, end int) ([]slack.AttachmentActionOption, error) {
	if begin >= end {
		return []slack.AttachmentActionOption{}, fmt.Errorf("begin must be smaller than end!")
	}

	attachemntOptions := make([]slack.AttachmentActionOption, end-begin+1)
	for i := range attachemntOptions {
		attachemntOptions[i] = slack.AttachmentActionOption{
			Text:  fmt.Sprintf("%d", i),
			Value: fmt.Sprintf("%d", i),
		}
	}

	return attachemntOptions, nil
}

func CancelButton(text string) slack.AttachmentAction {
	return slack.AttachmentAction{
		Name:  CANCEL_VAL,
		Text:  text,
		Type:  "button",
		Style: "danger",
	}
}

func ConfirmButton(text string) slack.AttachmentAction {
	return slack.AttachmentAction{
		Name:  CONFIRM_VAL,
		Text:  text,
		Type:  "button",
		Style: "primary",
	}
}

func Select(
	options []slack.AttachmentActionOption,
	text string,
	val string,
) slack.AttachmentAction {
	return slack.AttachmentAction{
		Name:    val,
		Text:    text,
		Type:    "select",
		Options: options,
	}
}
