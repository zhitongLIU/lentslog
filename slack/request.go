// Package ge slack

package slack

import (
	"errors"
	"fmt"
	slackSDK "github.com/nlopes/slack"
	"net/http"
	"os"
	"time"
)

type Request struct {
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

func NewRequest(r *http.Request) (*Request, error) {
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

	return &Request{
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

func FindDirectMessageDestinationUser(r *Request) (*slackSDK.User, error) {
	api := slackSDK.New(os.Getenv("SLACK_API_TOKEN"))
	imChannels, err := api.GetIMChannels()
	if err != nil {
		return nil, err
	}

	var userId string
	for _, imChannel := range imChannels {
		if imChannel.ID == r.ChannelId {
			userId = imChannel.User
		}
	}

	fmt.Println(userId)
	if userId == "" {
		return nil, errors.New("user id not found")
	}

	user, err := api.GetUserInfo(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
