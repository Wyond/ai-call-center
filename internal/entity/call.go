package entity

type Call struct {
	Id           string `json:"id" db:"id"`
	AudioInPath  string `json:"audio_in_path"`  // Путь к исходному аудио
	AudioOutPath string `json:"audio_out_path"` // Путь к ответу TTS
	Text         string `json:"text"`           // Транскрипция STT
	Response     string `json:"response"`       // Ответ NLP
}
