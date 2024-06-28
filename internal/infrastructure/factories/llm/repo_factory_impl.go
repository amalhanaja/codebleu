package llm

import (
	domainLlm "codebleu/internal/domain/llm"
	"codebleu/internal/infrastructure/thirdparty/gemini"
	"codebleu/pkg/env"
	"fmt"
	"strings"
)

func NewRepository(model string) (domainLlm.Repository, error) {
	if strings.HasPrefix(model, "gemini") {
		return gemini.NewClient(model, env.MustString("GEMINI_API_KEY")), nil
	}
	return nil, fmt.Errorf("invalid model: %s", model)
}
