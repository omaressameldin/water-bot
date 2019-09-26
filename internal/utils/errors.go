package utils

import (
	"fmt"
	"net/http"

	"github.com/shomali11/slacker"
)

func ReplyWithError(e error, m string, response slacker.ResponseWriter) {
	if e != nil {
		response.ReportError(fmt.Errorf("%s error: %s", m, e.Error()))
	}
}

func HttpError(e error, m string, w http.ResponseWriter) {
	if e != nil {
		w.Write([]byte(fmt.Sprintf("%s error: %s", m, e.Error())))
	}
}
