package main

import (
    mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Frigatier struct {
	config    *Config
	client    mqtt.Client
	eventsMap map[string]bool
}

func NewFrigatier() *Frigatier {
	return &Frigatier{}
}
