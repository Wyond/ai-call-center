package vosk

import "testing"

func TestTranscribe(t *testing.T) {
	s, err := New("vosk-model-ru")
	if err != nil {
		t.Fatal("cant set vosk model err:%v", err)
	}

	text, err := s.Transcribe("/home/vboxuser/ai-call-center/pkg/stt/vosk-api/go/example/phoneOff.wav")
	if err != nil {
		t.Fatal("cant transcribe audio err:%v", err)
	}

	t.Log("Transcribe text: ", text)
}
