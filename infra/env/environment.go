package main

import (
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type env struct {
	SlackToken string `envconfig:"SLACK_TOKEN"`
}

var (
	instance *env
	once     sync.Once
)

func GetEnv() *env{
	once.Do(func() {
		instance = &env{}
		envconfig.Process("", instance)
	})
	return instance
}