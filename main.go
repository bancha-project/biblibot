package main

import (
	"fmt"
	"github.com/bancha-project/biblibot/infra/env"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"log"
)

func main() {
	// .envから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Slack
	slackApi := slack.New(env.GetEnv().SlackToken)
	rtm := slackApi.NewRTM()
	go rtm.ManageConnection()

	// イベントを取得する
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			text := ev.Msg.Text
			var message string
			if text == "犬" {
				message = "🐶🐶🐶"
			} else if text == "猫" {
				message = "🐱😸🙀"
			} else if text == "kato" {
				message = "💢💢💢"
			}else {
				message = fmt.Sprintf("<@%v> hello!", ev.Msg.User)

			}
			rtm.SendMessage(rtm.NewOutgoingMessage(message, ev.Channel))
		}
	}
}
