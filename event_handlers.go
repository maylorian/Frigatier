package Frigatier

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

func (f *Frigatier) EventHandler(client mqtt.Client, msg mqtt.Message) {
	det := &Detection{}
	err := json.Unmarshal(msg.Payload(), det)
	if err != nil {
		log.Fatalf("Error reading detection from MQTT: %s", err)
	}
	if det.BeforeDetection.FalsePositive {
		return
	}
	eventId := det.BeforeDetection.Id
	_, ok := f.eventsMap[eventId]
	if !ok {
		f.eventsMap[eventId] = true
	} else {
		f.processAlreadySeenEvent(det)
		return
	}

	f.processNewEvent(det)
}

func (f *Frigatier) processAlreadySeenEvent(det *Detection) {
	beforeDetection := det.BeforeDetection
	if beforeDetection.EndTime != nil {
		delete(f.eventsMap, beforeDetection.Id)
	}
}

func (f *Frigatier) processNewEvent(msg *Detection) {
	slackEnabled := f.config.Messengers.Slack.Enabled
	if slackEnabled {
		slack := NewSlackClient(f.config.Messengers.Slack)
		f.notify(slack, msg)
	}
}

func (f *Frigatier) handlePostMessageActions(image string) {
	err := os.Remove(image)
	if err != nil {
		log.Fatalf("Failed to remove file: %s. Please delete manually.", image)
	}
}
