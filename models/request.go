// Package ge slack

package models

import (
	"net/http"
	"time"
)

type SlackRequest struct {
	id          int64
	TeamId      string
	TeamDomain  string
	UserId      string
	UserName    string
	Token       string
	ChannelId   string
	ChannelName string
	TriggerId   string
	Text        string
	Time        string
}

func NewSlackRequest(r *http.Request) (*SlackRequest, error) {
	r.ParseForm()
	err := r.ParseForm()

	if err != nil {
		// fmt.Fprintf(w, err.Error())
		return nil, err
	}

	token := r.Form.Get("token")
	text := r.Form.Get("text")
	channelId := r.Form.Get("channel_id")
	channelName := r.Form.Get("channel_name")
	userId := r.Form.Get("user_id")
	userName := r.Form.Get("user_name")
	teamId := r.Form.Get("team_id")
	teamDomain := r.Form.Get("team_domain")
	triggerId := r.Form.Get("trigger_id")

	timestamp := r.Header.Get("X-Slack-Request-Timestamp")
	requestTime, _ := time.Parse(time.UnixDate, timestamp)

	return &SlackRequest{
		Text:        text,
		Token:       token,
		TeamId:      teamId,
		TeamDomain:  teamDomain,
		UserId:      userId,
		UserName:    userName,
		ChannelId:   channelId,
		ChannelName: channelName,
		TriggerId:   triggerId,
		Time:        requestTime.String(),
	}, nil
}
