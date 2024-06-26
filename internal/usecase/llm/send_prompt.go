package llm

import (
	domain "codebleu/internal/domain/llm"
	"codebleu/internal/usecase"
	"context"
)

type sendPromptUseCase struct {
	repository domain.Repository
}

// Invoke implements usecase.UseCase.
func (s *sendPromptUseCase) Invoke(ctx context.Context, input string) (string, error) {
	return s.repository.SendPrompt(ctx, input)
}

func SendPromptUseCase(
	repository domain.Repository,
) usecase.UseCase[string, string] {
	return &sendPromptUseCase{
		repository: repository,
	}
}
