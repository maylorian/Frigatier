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
	config SlackConfig
}

func NewSlackClient(s SlackConfig) *Slack {
	return &Slack{
		client: slack.New(s.Token),
		config: s,
	}
}

func (s *Slack) Name() string {
	return "Slack"
}

func (s *Slack) IsEnabled() bool {
	return s.config.Enabled
}

func (s *Slack) Notify(msg *Detection, image string) error {
	message := fmt.Sprintf("A %s was detected by %s", msg.BeforeDetection.Label, msg.BeforeDetection.Camera)
	params := slack.FileUploadParameters{
		Channels:       []string{s.config.Channel},
		InitialComment: message,
		File:           image,
	}
	_, err := s.client.UploadFile(params)
	if err != nil {
		return err
	}
	return nil
}
