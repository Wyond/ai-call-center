package call

import (
	"ai-call-center/internal/entity"
	"context"
	"fmt"
	"log/slog"
)

type CallUseCase struct {
	repo CallRepository
	nlp  NLPService
	stt  STTService
	tts  TTSService
}

func NewCallUse(repo CallRepository, nlp NLPService, stt STTService, tts TTSService) *CallUseCase {
	return &CallUseCase{repo, nlp, stt, tts}
}

func (cu *CallUseCase) Process(ctx context.Context, audioPath string) error {
	//STT
	text, err := cu.stt.Transcribe(audioPath)
	if err != nil {
		slog.Error("Cant transcribe audio, err:", err)
		return fmt.Errorf("Cant transcribe audio, err:", err)
	}

	//NLP
	response, err := cu.nlp.Generate(text)
	if err != nil {
		return fmt.Errorf("cant load promt to generate, err:", err)
	}

	//TTS
	audioOut, err := cu.tts.Synthesize(response)
	if err != nil {
		return fmt.Errorf("cant synthesize our final text, err:", err)
	}

	call := &entity.Call{
		AudioInPath:  audioPath,
		AudioOutPath: audioOut,
		Text:         text,
		Response:     response,
	}
	return cu.repo.Save(ctx, call)
}
