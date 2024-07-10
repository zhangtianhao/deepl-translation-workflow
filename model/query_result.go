package model

// Result {"translations":[{"detected_source_language":"EN","text":"你好，世界"}]}
type Result struct {
	Translations []Translation `json:"translations"`
}
type Translation struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}
