package call

import (
	"ai-call-center/internal/entity"
	"context"
)

type CallRepository interface {
	Save(ctx context.Context, c *entity.Call) error
	GetById(ctx context.Context, id string) (*entity.Call, error)
}

type STTService interface {
	Transcribe(audioPath string) (string, error)
}

type NLPService interface {
	Generate(promt string) (string, error)
}

type TTSService interface {
	Synthesize(text string) (string, error)
}
