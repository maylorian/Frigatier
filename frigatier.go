package Frigatier

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Frigatier struct {
	config    *Config
	client    mqtt.Client
	eventsMap map[string]bool
}

func NewFrigatier() *Frigatier {
	return &Frigatier{}
}

func (f *Frigatier) Run() {
	f.sanityChecks()
	f.parseConfig()
	f.checkConfig()
	f.createMaps()
	f.createClient()
	f.subscribeToTopics()
	f.runForever()
}

func (f *Frigatier) sanityChecks() {
	_, err := os.Stat("config.yml")
	if errors.Is(err, os.ErrNotExist) {
		log.Fatalln("Could not find config.yml.")
	}
}

func (f *Frigatier) parseConfig() {
	b, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error during reading config file: %s", err.Error())
	}
	c := NewConfig()
	if err := yaml.Unmarshal(b, c); err != nil {
		log.Fatalf("Error persing config file: %s", err.Error())
	}
	f.config = c
}

func (f *Frigatier) checkConfig() {
	if f.config.Frigate.Host == "" {
		log.Fatalln("Please specific the frigate host.")
	}
	if f.config.Frigate.Port == 0 {
		log.Fatalln("Please specific the frigate port.")
	}
	if f.config.Mqtt.Host == "" {
		log.Fatalln("Please specific the mqtt host.")
	}
	if f.config.Mqtt.Port == 0 {
		log.Fatalln("Please specific the MQTT port.")
	}
}

func (f *Frigatier) createMaps() {
	f.eventsMap = make(map[string]bool)
}

func (f *Frigatier) createClient() {
	mqttConfig := f.config.Mqtt
	connectionString := fmt.Sprintf("tcp://%s:%d", mqttConfig.Host, mqttConfig.Port)
	opts := mqtt.NewClientOptions().AddBroker(connectionString).SetClientID("frigatier")
	opts.SetDefaultPublishHandler(f.EventHandler)
	if mqttConfig.User != "" {
		opts.SetUsername(mqttConfig.User)
	}
	if mqttConfig.Password != "" {
		opts.SetPassword(mqttConfig.Password)
	}
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	f.client = c
}

func (f *Frigatier) subscribeToTopics() {
	mqttConfig := f.config.Mqtt
	topicPrefix := mqttConfig.TopicPrefix
	if topicPrefix == "" {
		topicPrefix = "frigate"
	}
	eventsTopic := fmt.Sprintf("%s/events", topicPrefix)
	if token := f.client.Subscribe(eventsTopic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf(token.Error().Error())
	}
}

func (f *Frigatier) runForever() {
	for {
		select {}
	}
}
