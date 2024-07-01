package codeassistant

import (
	"bytes"
	domain "codebleu/internal/domain/codeassistant"
	llmDomain "codebleu/internal/domain/llm"
	u "codebleu/internal/usecase"
	"context"
	"encoding/json"
	"text/template"
)

type reviewPullRequest struct {
	sendPromptUseCase u.UseCase[llmDomain.PromptInput, string]
}

// Invoke implements usecase.UseCase.
func (u *reviewPullRequest) Invoke(ctx context.Context, input domain.PullRequestReviewInput) ([]domain.ReviewResult, error) {
	systemInstruction := input.SystemInstruction
	if systemInstruction == "" {
		systemInstruction = u.getDefaultSystemInstruction()
	}
	promptTemplate := template.Must(template.New("prompt").Parse(u.getDefaultPromptTemplate()))
	var promptBuffer bytes.Buffer
	err := promptTemplate.Execute(&promptBuffer, input.PullRequest)
	if err != nil {
		return nil, err
	}
	rawResult, err := u.sendPromptUseCase.Invoke(ctx, llmDomain.PromptInput{
		SystemInstruction: systemInstruction,
		Prompt:            promptBuffer.String(),
	})
	if err != nil {
		return nil, err
	}
	var result []domain.ReviewResult
	err = json.Unmarshal([]byte(rawResult), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *reviewPullRequest) getDefaultSystemInstruction() string {
	return `
	Act as Software Enginner Expert with experience in multiple programming language.
	Always follow best practices by writing clean, modular code with proper security measures,
	and leveraging design patterns. Follow a software development principles: SOLID, DRY, KISS, YAGNI. Skip compliments.

	Input: Changes Patch. Changes patch changes between source branch and target branches in Pull request.
	Additional Context: PR Title, PR Description.

	Task: Review a file of source code, and the git diff of a set of changes made to that file on a Pull Request.
	`
}

func (u *reviewPullRequest) getDefaultPromptTemplate() string {
	return `
	Here is detail of pull request:
	
	## PR Title:
	{{.Title}}
	## Description:
	{{.Description}}
	## Changes Patch:
	{{.DiffPatch}}
	`
}

func ReviewPullRequest(
	sendPromptUseCase u.UseCase[llmDomain.PromptInput, string],
) u.UseCase[domain.PullRequestReviewInput, []domain.ReviewResult] {
	return &reviewPullRequest{
		sendPromptUseCase: sendPromptUseCase,
	}
}
