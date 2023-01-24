package Frigatier

type Config struct {
	Frigate    Frigate                `yaml:"frigate"`
	Mqtt       Mqtt                   `yaml:"mqtt"`
	Messengers NotificationMessengers `yaml:"messengers"`
}

type Frigate struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Mqtt struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	TopicPrefix string `yaml:"topic_prefix"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
}

type NotificationMessengers struct {
	Slack   SlackConfig   `yaml:"slack"`
	Discord DiscordConfig `yaml:"discord"`
}

func NewConfig() *Config {
	return &Config{}
}
