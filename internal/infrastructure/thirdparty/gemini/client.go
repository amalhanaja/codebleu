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
func (c *client) SendPrompt(ctx context.Context, prompt string) (string, error) {
	genaiClient, err := c.getClient(ctx)
	if err != nil {
		return "", err
	}
	defer genaiClient.Close()
	model := genaiClient.GenerativeModel(c.model)
	chat := model.StartChat()
	chat.History = []*genai.Content{}
	resp, err := chat.SendMessage(ctx, genai.Text(prompt))
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
