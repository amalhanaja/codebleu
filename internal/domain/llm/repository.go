package llm

import "context"

type Repository interface {
	SendPrompt(ctx context.Context, prompt string) (string, error)
}
