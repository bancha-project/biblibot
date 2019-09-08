package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"

	"github.com/bancha-project/biblibot/infra/env"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"gopkg.in/yaml.v2"
)

type ReplyDic struct {
	Keyword string
	Replies []string
}

func main() {
	// .envから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}

	buf, err := ioutil.ReadFile("./infra/data/reply_dic.yaml")
	if err != nil {
		log.Fatal("Error loading replies file", err)
		return
	}

	replyDics := []ReplyDic{}
	err = yaml.Unmarshal(buf, &replyDics)
	if err != nil {
		log.Fatal("Error yaml unmarshaling", err)
		return
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

			// 辞書のキーワードにマッチする返信をランダムで返す
			for _, replyDic := range replyDics {
				if strings.Contains(strings.ToLower(text), strings.ToLower(replyDic.Keyword)) {
					replies := replyDic.Replies
					message = fmt.Sprintf("<@%v> %v", ev.Msg.User, replies[rand.Intn(len(replies))])
					break
				}
			}

			rtm.SendMessage(rtm.NewOutgoingMessage(message, ev.Channel))
		}
	}
}
