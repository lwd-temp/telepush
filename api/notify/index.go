package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lwd-temp/telepush/internal/bot"
	"github.com/lwd-temp/telepush/internal/config"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userId := config.ReceiverID
	msg := params.Get("m")
	token := params.Get("t")

	if token != config.APIToken {
		fmt.Fprintf(w, "token error")
		return
	}

	if err := bot.NotifyTxtMessage(userId, msg); err != nil {
		fmt.Fprintf(w, "send message to %d failed!", userId)
		return
	}

	fmt.Fprintf(w, "send message to %d succes!", userId)
}
