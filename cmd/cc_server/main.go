package cc_server

import (
	"ai-call-center/pkg/stt/vosk"
	"log"
)

func main() {
	stt, err := vosk.New("models/vosk-model-small-ru")
	if err != nil {
		log.Fatal("Fai;ed to init Vosk, err: ", err)
	}
}
