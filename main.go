package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/YuyaBan/golang_SlackBot/domain"
	"github.com/nlopes/slack"
)

func main() {
	env := domain.NewEnviroment()

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

	rand.Seed(time.Now().UnixNano())
	emoji := []string{" :nahemohu: ", " :nahemohu_guruguru: ", " :nahemohu_scroll: ", " :nahemohu_suisui: "}
	ans := emoji[rand.Intn(4)]

	if strings.Contains(ev.Msg.Text, "妹") {
		rtm.SendMessage(rtm.NewOutgoingMessage(":imouto: が欲しいのか？ わかるぞ。", ev.Channel))
	} else if strings.Contains(ev.Msg.Text, "なへもふ") {
		rtm.SendMessage(rtm.NewOutgoingMessage(":nahemohu: がそんなに好き？", ev.Channel))
	} else if strings.Contains(ev.Msg.Text, "なへもふ") {
		rtm.SendMessage(rtm.NewOutgoingMessage("なへもふ占い！ 今日のお前は"+ans+"な日!!", ev.Channel))
	} else {
		rtm.SendMessage(rtm.NewOutgoingMessage("呼んだ？\n「妹」「なへもふ」「占い」に反応するぞ。", ev.Channel))
	}

	return nil
}
