package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// .envから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("SLACK_TOKEN")
	println(token)
}
