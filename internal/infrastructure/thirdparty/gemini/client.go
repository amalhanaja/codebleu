package gemini

import (
	domain "codebleu/internal/domain/llm"
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type client struct {
	apiKey string
	model  string
}

// SendPrompt implements llm.Repository.
func (c *client) SendPrompt(ctx context.Context, input domain.PromptInput) (string, error) {
	genaiClient, err := c.getClient(ctx)
	if err != nil {
		return "", err
	}
	systemOutputInstruction := `
		#IMPORTANT INSTRUCTION:\n
		Returns using this JSON scheme: [{path: <string>,comment_in_markdown: <string>}]
	`
	defer genaiClient.Close()
	model := genaiClient.GenerativeModel(c.model)
	model.GenerationConfig = genai.GenerationConfig{
		ResponseMIMEType: "application/json",
	}
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(input.SystemInstruction),
			genai.Text(systemOutputInstruction),
		},
	}
	resp, err := model.GenerateContent(ctx, genai.Text(input.Prompt))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%+v", resp.Candidates[0].Content.Parts[0]), nil
}

func (c *client) getClient(ctx context.Context) (*genai.Client, error) {
	return genai.NewClient(ctx, option.WithAPIKey(c.apiKey))
}

func NewClient(model string, apiKey string) domain.Repository {
	return &client{
		apiKey: apiKey,
		model:  model,
	}
}
