package main

import (
	"log"
	"os"
	"strings"

	"github.com/YuyaBan/golang_SlackBot/app/domain"
	"github.com/gin-gonic/gin"
	"github.com/nlopes/slack"
)

func main() {
	env := domain.NewEnviroment()

	// for heroku
	port := os.Getenv("PORT")
	router := gin.New()
	router.Run(":" + port)
	//

	rtm := slack.New(env.Token).NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := ValidateMessageEvent(ev, rtm, env); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// ValidateMessageEvent send Message
func ValidateMessageEvent(ev *slack.MessageEvent, rtm *slack.RTM, env *domain.Config) error {

	// Only response mention to bot. Ignore else.
	if !strings.HasPrefix(ev.Msg.Text, env.Botid) {
		log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return nil
	}

	if strings.Contains(ev.Msg.Text, "imouto") {
		rtm.SendMessage(rtm.NewOutgoingMessage(":imouto: が欲しいのか？ わかるぞ。", ev.Channel))
	} else if strings.Contains(ev.Msg.Text, "nahemohu") {
		rtm.SendMessage(rtm.NewOutgoingMessage(":nahemohu: がそんなに好き？", ev.Channel))
	} else {
		rtm.SendMessage(rtm.NewOutgoingMessage("呼んだ？。どうした？", ev.Channel))
	}

	return nil
}
