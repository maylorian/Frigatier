package Frigatier

import (
	"encoding/json"
	"fmt"
	"github.com/andreasavg/Frigatier/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

func (f *Frigatier) EventHandler(client mqtt.Client, msg mqtt.Message) {
	det := &Detection{}
	err := json.Unmarshal(msg.Payload(), det)
	utils.WarnIfErr(err, "Error reading detection from MQTT")
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
	image := f.getEventImage(msg.BeforeDetection.Id)
	for _, messenger := range f.enabledMessengers {
		err := messenger.Notify(msg, image)
		utils.WarnIfErr(err, "Failure to send message through slack.")
	}
	f.handlePostMessageActions(image)
}

func (f *Frigatier) handlePostMessageActions(image string) {
	err := os.Remove(image)
	utils.WarnIfErr(err, fmt.Sprintf("Failed to remove file: %s. Please delete manually.", image))
}
