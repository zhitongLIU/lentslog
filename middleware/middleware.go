package middleware

import (
	"fmt"
	"github.com/zhitongLIU/lentslog/slack"
	"net/http"
	"os"
)

func CertifyApp(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slackKey := os.Getenv("SLACK_APP_TOKEN")

		request, err := slack.NewRequest(r)

		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		if slackKey != request.Token {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Not right key"))
			return
		}
		handler.ServeHTTP(w, r)
	}
}
