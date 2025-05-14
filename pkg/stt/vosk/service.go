package vosk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	vosk "github.com/alphacep/vosk-api/go"
	//osk "stt/vosk-api/go"
)

type Service struct {
	model *vosk.VoskModel
}

func New(modelPath string) (*Service, error) {
	model, err := vosk.NewModel(modelPath)
	if err != nil {
		return nil, err
	}
	return &Service{model: model}, nil
}

func (s *Service) Transcribe(audioPath string) (string, error) {
	rec, err := vosk.NewRecognizer(s.model, 16000.0)
	if err != nil {
		return "", fmt.Errorf("cant Recognized vosk model, err:", err)
	}
	defer rec.Free()

	file, err := os.Open(audioPath)
	if err != nil {
		log.Fatal("cant open audio file, err:%v", err)
	}

	buf := make([]byte, 4096)

	for {
		_, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}

		if rec.AcceptWaveform(buf) != 0 {
			fmt.Println(rec.Result())
		}
	}

	var jres map[string]interface{}
	json.Unmarshal([]byte(rec.FinalResult()), &jres)
	fmt.Println(jres["text"])

	return rec.FinalResult(), nil
}
