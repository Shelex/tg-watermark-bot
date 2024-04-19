package main

import (
	"Shelex/tg-watermark-bot/locale"
	"fmt"
	"os"
	"path/filepath"
)

func GetWatermarkPath(userID int64) (string, error) {
	var file string
	pattern := fmt.Sprintf("%d-watermark.*", userID)
	matches, err := filepath.Glob(pattern)

	if err != nil {
		return "", locale.Error("failed_watermark_lookup")
	}

	if len(matches) > 0 {
		file = matches[0]
	}

	if file == "" {
		return "", locale.Error("failed_watermark_lookup")
	}

	return file, nil
}

func WatermarkExists(userID int64) bool {
	_, err := GetWatermarkPath(userID)
	return err == nil
}

func RemoveWatermark(userID int64) {
	watermark, err := GetWatermarkPath(userID)
	if err != nil {
		return
	}

	os.Remove(watermark)
}
