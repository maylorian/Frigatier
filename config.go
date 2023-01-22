package main

type Config struct {
	Mqtt       Mqtt       `yaml:"mqtt"`
	Messengers Messengers `yaml:"messengers"`
}
type Mqtt struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	TopicPrefix string `yaml:"topic_prefix"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
}

type Messengers struct {
	Slack SlackMessenger `yaml:"slack"`
}

func NewConfig() *Config {
	return &Config{}
}
