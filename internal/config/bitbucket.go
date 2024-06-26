package config

import "codebleu/pkg/env"

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

func NewBitbucketWebhookConfigFromEnv() BitbucketWebhookConfig {
	return BitbucketWebhookConfig{
		Secret: env.MustString(bitbucketWebhookSecret),
	}
}

func NewBitbucketRepositoryAccessConfigFromEnv() BitbucketRepositoryAccessConfig {
	return BitbucketRepositoryAccessConfig{
		Workspace:   env.MustString(bitbucketWorkspace),
		RepoSlug:    env.MustString(bitbucketRepoSlug),
		AccessToken: env.MustString(bitbucketAccessToken),
	}
}
