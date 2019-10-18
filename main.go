package main

import (
	"log"
	"strings"

	"github.com/YuyaBan/golang_SlackBot/domain"
	"github.com/nlopes/slack"
)

// type Config struct {
// 	API APIConfig
// }

// type APIConfig struct {
// 	Token string
// 	Botid string
// }

func main() {
	// var config Config

	// _, err := toml.DecodeFile("config.toml", &config)
	// if err != nil {
	// 	// Error Handling
	// }

	// api := slack.New(
	// 	config.API.Token,
	// )

	// fmt.Println(config.API.Token)
	// rtm := api.NewRTM()
	env := domain.NewEnviroment()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := config.ValidateMessageEvent(ev, rtm); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// ValidateMessageEvent send Message
func (config *Config) ValidateMessageEvent(ev *slack.MessageEvent, rtm *slack.RTM) error {
	// Only response in specific channel. Ignore else.

	// Only response mention to bot. Ignore else.
	log.Print(ev.Msg.Text)
	if !strings.HasPrefix(ev.Msg.Text, config.API.Botid) {
		//log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return nil
	}

	rtm.SendMessage(rtm.NewOutgoingMessage("メンション付いてるな。:imouto: くれるのか？", ev.Channel))
	return nil
}
