package attachments

import (
	"encoding/json"
	"fmt"

	"github.com/nlopes/slack"
)

type SelectedOption struct {
	Name string
	Val  string
}

func MarshalSelectedOptions(selected []SelectedOption) (string, error) {
	marshalledSelected, err := json.Marshal(selected)
	if err != nil {
		return "", err
	}

	return string(marshalledSelected), nil
}

func UnMarshalSelectedOptions(selected string) ([]SelectedOption, error) {
	var unMarshalledSelected []SelectedOption
	err := json.Unmarshal([]byte(selected), &unMarshalledSelected)
	if err != nil {
		return []SelectedOption{}, nil
	}

	return unMarshalledSelected, nil
}

func AddAnswer(answerName, answerVal string) ([]SelectedOption, error) {
	answers, err := UnMarshalSelectedOptions(
		answerName,
	)
	if err != nil {
		return []SelectedOption{}, err
	}

	answers[len(answers)-1].Val = answerVal

	return answers, nil
}

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
