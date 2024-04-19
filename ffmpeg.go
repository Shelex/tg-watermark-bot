package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertVideoToWatermarked(waterPath string, videoPath string, destPath string) error {
	inputStream := ffmpeg.Input(videoPath)

	watermarkStream := ffmpeg.Input(waterPath).
		Filter("scale", ffmpeg.Args{"250:-1"}).
		Filter("format", ffmpeg.Args{"rgba"}).
		// watermark opacity is set via "colorchannelmixer"
		Filter("colorchannelmixer", ffmpeg.Args{"aa=0.4"})

	err := ffmpeg.Filter([]*ffmpeg.Stream{inputStream, watermarkStream},
		"overlay", ffmpeg.Args{"(main_w-overlay_w)/2:(main_h-overlay_h)/2", "format=auto"}).
		// "yuv420p" fixes video display for android devices
		// "map" argument fixes missing audio
		Output(destPath, ffmpeg.KwArgs{"pix_fmt": "yuv420p", "map": "0:a:?"}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()

	return err
}
