package vosk

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// func Speex(inputPath string) (string, error) {
// 	outputPath := "cleanSPEEX_" + filepath.Base(inputPath)
// 	cmd := exec.Command("speexenc",
// 		"--enhance",
// 		"--quality", "8",
// 		inputPath,
// 		outputPath+".spx",
// 	)
// 	if err := cmd.Run(); err != nil {
// 		return "", fmt.Errorf("cant clean audio(Speex)err:%v", err)
// 	}
// 	return outputPath, nil

// }

func CleanAudio(inputPath string) (string, error) {
	outputPath := "clean_" + filepath.Base(inputPath)
	// cmd := exec.Command("ffmpeg",
	// 	"-i", inputPath,
	// 	//"-af", "arnndn=m=models/rnnoise/rnnoise.ru.fb16000.bin",
	// 	"-af", "highpass=f=200,lowpass=f=3000,afftdn=nf=-20",
	// 	"-ar", "16000",
	// 	"-ac", "1",
	// 	"-y",
	// 	outputPath,
	// )
	cmd := exec.Command("ffmpeg",
		"-i", inputPath,
		"-af",
		`
	highpass=f=80,
	lowpass=f=4000,
	afftdn=nf=-20,
	loudnorm=I=-16:TP=-1.5:LRA=11
		`,
		"-ar", "16000",
		"-ac", "1",
		"-y",
		outputPath,
	)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cant clean audio(ffmpeg)err:%v", err)
	}
	return outputPath, nil
}

func CleanAudioSox(inputPath string) (string, error) {
	outputPathSox := strings.TrimSuffix(inputPath, ".wav") + "_sox.wav"
	cmd := exec.Command("sox",
		inputPath,
		outputPathSox,
		"highpass", "80", //удаление низкочастотного шума
		"lowpass", "4000", //срез высоких шумов
		//"compand", "0.02,0.05", //динамическое сжатие
		"noisered", "noise.prof", "0.2", //подавление шума
		"norm", "-0.1", //нормализация
		"rate", "16000",
		"channels", "1",
		//"compand", "0.02", "0.05",
	)
	cmd.Stderr = os.Stderr
	return outputPathSox, cmd.Run()
}
