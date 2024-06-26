package httpserver

import (
	"codebleu/internal/config"
	"codebleu/pkg/env"
)

// AppConfig contains application configuration
type AppConfig struct {
	Port                            string
	BitbucketRepositoryAccessConfig config.BitbucketRepositoryAccessConfig
	BitbucketWebhookConfig          config.BitbucketWebhookConfig
	GeminiConfig                    config.GeminiConfig
}

func NewConfigFromEnv() AppConfig {
	return AppConfig{
		Port:                            env.MustString(config.EnvServerPort),
		BitbucketRepositoryAccessConfig: config.NewBitbucketRepositoryAccessConfigFromEnv(),
		BitbucketWebhookConfig:          config.NewBitbucketWebhookConfigFromEnv(),
		GeminiConfig:                    config.NewGeminiConfigFromEnv(),
	}
}
