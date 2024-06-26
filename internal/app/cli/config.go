package cli

import (
	"fmt"
	"os"
)

// AppConfig contains application configuration
type AppConfig struct {
	BitbucketRepositoryAccessConfig *BitbucketRepositoryAccessConfig
	GeminiConfig                    *GeminiConfig
}

// BitbucketRepositoryAccessConfig contains Config to access bitbucket Repository
type BitbucketRepositoryAccessConfig struct {
	Workspace   string
	RepoSlug    string
	AccessToken string
}

// BitbucketWebhookConfig contains Config to configure BitbucketWebhook
type BitbucketWebhookConfig struct {
	Secret string
}

// GeminiConfig contains config to configure Gemini
type GeminiConfig struct {
	ApiKey string
}

func mustEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		panic(fmt.Sprintf("Env key %s not found", key))
	}
	return env
}

func NewConfigFromEnv() *AppConfig {
	cfg := &AppConfig{}
	cfg.overrideBitbucketRepositoryAccessConfig()
	cfg.overrideGeminiConfig()
	return cfg
}

func (cfg *AppConfig) overrideGeminiConfig() {
	cfg.GeminiConfig = &GeminiConfig{}
	cfg.GeminiConfig.ApiKey = mustEnv("GEMINI_API_KEY")
}

func (cfg *AppConfig) overrideBitbucketRepositoryAccessConfig() {
	cfg.BitbucketRepositoryAccessConfig = &BitbucketRepositoryAccessConfig{}
	cfg.BitbucketRepositoryAccessConfig.AccessToken = mustEnv("BITBUCKET_ACCESS_TOKEN")
	cfg.BitbucketRepositoryAccessConfig.RepoSlug = mustEnv("BITBUCKET_REPOSITORY_SLUG")
	cfg.BitbucketRepositoryAccessConfig.Workspace = mustEnv("BITBUCKET_WORKSPACE")
}
