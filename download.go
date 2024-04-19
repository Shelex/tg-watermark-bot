package main

import (
	"Shelex/tg-watermark-bot/locale"
	"errors"
	"io"
	"net/http"
	"os"
)

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return locale.WithError("failed_download_file", errors.New(resp.Status))
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
