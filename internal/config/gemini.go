package config

import "codebleu/pkg/env"

// GeminiConfig contains config to configure Gemini
type GeminiConfig struct {
	ApiKey string
}

func NewGeminiConfigFromEnv() GeminiConfig {
	return GeminiConfig{
		ApiKey: env.MustString(geminiApiKey),
	}
}
