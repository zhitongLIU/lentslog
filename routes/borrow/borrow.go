package borrow

import (
	"errors"
	"fmt"
	"github.com/zhitongLIU/lentslog/slack"
	"github.com/zhitongLIU/lentslog/transaction"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	request, err := slack.NewRequest(r)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Printf("%s = %s\n", "ChannelId", request.ChannelId)
	fmt.Printf("%s = %s\n", "Token", request.Token)
	fmt.Printf("%s = %s\n", "TriggerId", request.TriggerId)
	fmt.Printf("%s = %s\n", "TeamId", request.TeamId)
	fmt.Printf("%s = %s\n", "TeamDomain", request.TeamDomain)
	fmt.Printf("%s = %s\n", "UserId", request.UserId)
	fmt.Printf("%s = %s\n", "Text", request.Text)
	fmt.Printf("%s = %s\n", "ChannelName", request.ChannelName)
	fmt.Printf("%s = %s\n", "Time", request.Time)
	fmt.Printf("%s = %s\n", "UserName", request.UserName)

	dump, err := httputil.DumpRequest(r, true)
	fmt.Fprintf(w, string(dump))

	for k, v := range r.Form {
		fmt.Fprintf(w, "%s = %s\n", k, v)
	}

	u, err := slack.FindDirectMessageDestinationUser(request)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, u.Name)

	text := request.Text
	sum, _ := strconv.ParseFloat(text, 64)

	err = valid_borrow(sum)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Printf("%T, %v\n", sum, sum)
	transaction, err := transaction.New(sum, string(request.UserId), string(u.ID))
	fmt.Println("transaction %s", transaction)
}

func valid_borrow(sum float64) error {
	err_msg := ""
	if sum <= 0 {
		err_msg += "borrow should be positive! "
	}

	if err_msg != "" {
		return errors.New(err_msg)
	}
	return nil
}
