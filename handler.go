package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"Shelex/tg-watermark-bot/env"
	"Shelex/tg-watermark-bot/locale"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgApi struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
	Config *env.Config
}

func NewAPI(bot *tgbotapi.BotAPI, update tgbotapi.Update, config *env.Config) TgApi {
	return TgApi{
		Bot:    bot,
		Update: update,
		Config: config,
	}
}

func (api *TgApi) GetUserID() (int64, error) {
	from := api.Update.Message.From

	if from == nil {
		return 0, locale.Error("unknown_tg_user_id")
	}

	return from.ID, nil
}

func (api *TgApi) SendTextMessage(msg string) {
	message := tgbotapi.NewMessage(api.Update.Message.Chat.ID, "")
	message.Text = msg

	if _, err := api.Bot.Send(message); err != nil {
		log.Panic(err)
	}
}

func (api *TgApi) HandleWatermarkAttachment(userID int64) error {
	api.SendTextMessage(locale.Translate("saving_watermark"))

	fileID := getFileID(api.Update)

	if fileID == "" {
		return locale.Error("unknown_file_id")
	}

	file, err := api.Bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return locale.Error("tg_get_file_failed")
	}

	link := file.Link(api.Config.TgToken)

	log.Println(link)

	filenameFromLink := path.Base(link)
	ext := filepath.Ext(filenameFromLink)

	filename := fmt.Sprintf("%d-watermark%s", userID, ext)

	if err := downloadFile(filename, link); err != nil {
		return locale.WithError("download_watermark_failed", err)
	}

	api.SendTextMessage(locale.Translate("watermark_saved"))
	return nil
}

func (api *TgApi) HandleVideoAttachment(userID int64) error {
	hasWatermark := WatermarkExists(userID)

	if !hasWatermark {
		return locale.Error("no_watermark")
	}

	api.SendTextMessage(locale.Translate("video_detected"))

	video := api.Update.Message.Video

	file, err := api.Bot.GetFile(tgbotapi.FileConfig{FileID: video.FileID})

	if err != nil {
		return locale.WithError("tg_get_video_failed", err)
	}

	link := file.Link(api.Config.TgToken)
	ext := filepath.Ext(filepath.Base(link))
	filename := fmt.Sprintf("%d-video%s", userID, ext)

	if err := downloadFile(filename, link); err != nil {
		return locale.WithError("download_video_failed", err)
	}

	api.SendTextMessage(locale.Translate("video_downloaded"))

	watermarkPath, err := GetWatermarkPath(userID)
	if err != nil {
		return locale.WithError("read_watermark_failed", err)
	}

	destPath := fmt.Sprintf("%d-converted%s", userID, ext)

	api.SendTextMessage(locale.Translate("start_overlaying_video"))

	if err := ConvertVideoToWatermarked(string(watermarkPath), filename, destPath); err != nil {
		return locale.WithError("video_overlay_failed", err)
	}

	watermarked := tgbotapi.NewVideo(api.Update.Message.Chat.ID, tgbotapi.FilePath(destPath))
	watermarked.ReplyToMessageID = api.Update.Message.MessageID

	msg, err := api.Bot.Send(watermarked)

	if err != nil {
		return locale.WithError("upload_overlay_video_failed", err)
	}

	log.Printf("video sent: %+v\n", msg)

	os.Remove(filename)
	os.Remove(destPath)
	return nil
}
