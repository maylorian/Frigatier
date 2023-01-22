package Frigatier

import (
	"fmt"
	"github.com/slack-go/slack"
)

type SlackConfig struct {
	Enabled      bool   `yaml:"enabled"`
	SendSnapshot bool   `yaml:"send_snapshot"`
	Token        string `yaml:"token"`
	Channel      string `yaml:"channel"`
}

type Slack struct {
	client *slack.Client
}

func NewSlackClient(s SlackConfig) *slack.Client {
	return slack.New(s.Token)
}

func (f *Frigatier) notify(s *slack.Client, msg *Detection) error {
	message := fmt.Sprintf("A %s was detected by %s", msg.BeforeDetection.Label, msg.BeforeDetection.Camera)
	image := f.getEventImage(msg.BeforeDetection.Id)
	params := slack.FileUploadParameters{
		Channels:       []string{f.config.Messengers.Slack.Channel},
		InitialComment: message,
		File:           image,
	}
	_, err := s.UploadFile(params)
	if err != nil {
		return err
	}
	f.handlePostMessageActions(image)
	return nil
}
