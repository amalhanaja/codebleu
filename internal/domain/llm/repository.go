package llm

import "context"

type Repository interface {
	SendPrompt(ctx context.Context, input PromptInput) (string, error)
}
