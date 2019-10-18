package main

import (
	"log"
	"strings"

	"github.com/YuyaBan/golang_SlackBot/domain"
	"github.com/nlopes/slack"
)

type Env struct {
	Config Config
	rtm    *slack.RTM
}

type Config struct {
	Token string
	Botid string
}

func main() {
    var env Env
    env.Config := domain.NewEnviroment()

	env.rtm := slack.New(env.Config.Token).NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := env.ValidateMessageEvent(ev, rtm); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// ValidateMessageEvent send Message
func (env *Env) ValidateMessageEvent(ev *slack.MessageEvent, rtm *slack.RTM) error {
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
