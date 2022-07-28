package config

import "os"

var (
	BotToken = os.Getenv("telegram_bot_token")
	APIToken = os.Getenv("api_token")
)
