package env

import (
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type env struct {
	SlackToken string `envconfig:"SLACK_TOKEN"`
	BotId      string `envconfig:"BOT_ID"`
	ChannelId  string `envconfig:"CHANNEL_ID"`
}

var (
	instance *env
	once     sync.Once
)

func GetEnv() *env {
	once.Do(func() {
		instance = &env{}
		envconfig.Process("", instance)
	})
	return instance
}
