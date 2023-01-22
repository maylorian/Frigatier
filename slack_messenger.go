package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
)

type SlackMessenger struct {
	Enabled bool   `yaml:"enabled"`
	Token   string `yaml:"token"`
	Channel string `yaml:"channel"`
}

func NewSlackClient(s SlackMessenger) *slack.Client {
	return slack.New(s.Token)
}

func (f *Frigatier) notifySlack(s *slack.Client, msg *Detection) {
	message := fmt.Sprintf("A %s was detected by %s", msg.BeforeDetection.Label, msg.BeforeDetection.Camera)
	_, _, err := s.PostMessage(f.config.Messengers.Slack.Channel, slack.MsgOptionText(message, false))
	if err != nil {
		log.Println("Failed to send message to slack.")
	}
}
