package Frigatier

import "github.com/slack-go/slack"

type Messenger interface {
	notify(*slack.Client, *Detection) error
}
