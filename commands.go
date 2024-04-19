package main

import (
	"fmt"

	"Shelex/tg-watermark-bot/env"
	"Shelex/tg-watermark-bot/locale"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isPhoto(update tgbotapi.Update) bool {
	photo := update.Message.Photo
	return len(photo) > 0
}

func isVideo(update tgbotapi.Update) bool {
	video := update.Message.Video
	return video != nil
}

func isDocument(update tgbotapi.Update) bool {
	doc := update.Message.Document
	return doc != nil
}

func getFileID(update tgbotapi.Update) string {
	if isDocument(update) {
		return update.Message.Document.FileID
	}
	if isPhoto(update) {
		photos := update.Message.Photo
		lastPhoto := photos[len(photos)-1]
		return lastPhoto.FileID
	}

	return ""
}

func HandleTgUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *env.Config) {
	if update.Message == nil {
		return
	}

	api := NewAPI(bot, update, config)

	userID, err := api.GetUserID()

	if err != nil {
		api.SendTextMessage(err.Error())
		return
	}

	switch update.Message.Command() {
	case "help":
		api.SendTextMessage(HelpText())
	case "start":
		api.SendTextMessage(fmt.Sprintf("%s, %s", Greet(update.Message.From.FirstName), HelpText()))
	case "status":
		api.SendTextMessage(locale.Translate("status_ok"))
	case "clear_watermark":
		RemoveWatermark(userID)
		api.SendTextMessage(locale.Translate("removed_watermark"))
	default:
		if isPhoto(update) || isDocument(update) {
			if err := api.HandleWatermarkAttachment(userID); err != nil {
				api.SendTextMessage(err.Error())
			}
			return
		}

		if isVideo(update) {
			if err := api.HandleVideoAttachment(userID); err != nil {
				api.SendTextMessage(err.Error())
			}
			return
		}

		api.SendTextMessage(locale.Translate("unknown_command"))
	}
}
