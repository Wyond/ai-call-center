package vosk

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func CleanAudio(inputPath string) (string, error) {
	outputPath := "clean_" + filepath.Base(inputPath)
	cmd := exec.Command("ffmpeg",
		"-i", inputPath,
		"-af", "highpass=f=200,lowpass=f=3000,afftdn=nf=-20",
		"-ar", "16000",
		"-ac", "1",
		outputPath,
	)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cant clean audio(ffmpeg)err:%v", err)
	}
	return outputPath, nil
}

func CleanAudioSox(inputPath string) (string, error) {
	outputPathSox := "clean_sox.wav"
	cmd := exec.Command("sox",
		inputPath,
		outputPathSox,
		"noisered", "noise.prof", "0.2",
		"compand", "0.02", "0.05",
	)
	return outputPathSox, cmd.Run()
}
