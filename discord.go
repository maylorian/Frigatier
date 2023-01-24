package Frigatier

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

type DiscordConfig struct {
	Enabled      bool   `yaml:"enabled"`
	SendSnapshot bool   `yaml:"send_snapshot"`
	Token        string `yaml:"token"`
	Channel      string `yaml:"channel"`
}
type Discord struct {
	config DiscordConfig
}

func NewDiscordMessenger(d DiscordConfig) *Discord {
	return &Discord{
		config: d,
	}
}
func (d *Discord) Notify(msg *Detection, image string) error {
	message := fmt.Sprintf("A %s was detected by %s", msg.BeforeDetection.Label, msg.BeforeDetection.Camera)
	discord, err := discordgo.New("Bot " + d.config.Token)
	if err != nil {
		log.Println("Failed to create a discord client.")
	}
	file, err := os.Open(image)
	if err != nil {
		log.Println("Failed to read image. Sending plain message.")
		discord.ChannelMessageSend(d.config.Channel, message)
	}
	discord.ChannelFileSend(d.config.Channel, message, file)
	return nil
}

func (d *Discord) Name() string {
	return "Discord"
}

func (d *Discord) IsEnabled() bool {
	return d.config.Enabled
}
