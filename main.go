package main

import (
	"log"

	"Shelex/tg-watermark-bot/env"
	"Shelex/tg-watermark-bot/locale"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := env.ReadEnv()

	locale.Register(config)

	bot, err := tgbotapi.NewBotAPI(config.TgToken)
	if err != nil {
		translated := locale.WithError("error_initialize", err)
		log.Panic(translated.Error())
	}

	if config.Environment == "dev" {
		bot.Debug = true
	}

	log.Printf("%s %s", locale.Translate("authorized_for"), bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		HandleTgUpdate(bot, update, config)
		continue
	}
}
