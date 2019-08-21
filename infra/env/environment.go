package env

import (
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type env struct {
	SlackToken string `envconfig:"SLACK_TOKEN"`
	BotUser    string `envconfig:"BOT_USER"`
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
