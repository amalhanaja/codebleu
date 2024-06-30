package codeassistant

import (
	"bytes"
	domain "codebleu/internal/domain/codeassistant"
	llmDomain "codebleu/internal/domain/llm"
	u "codebleu/internal/usecase"
	"context"
	"text/template"
)

type reviewPullRequest struct {
	sendPromptUseCase u.UseCase[llmDomain.PromptInput, string]
}

// Invoke implements usecase.UseCase.
func (u *reviewPullRequest) Invoke(ctx context.Context, input domain.PullRequestReviewInput) (string, error) {
	systemInstruction := input.SystemInstruction
	if systemInstruction == "" {
		systemInstruction = u.getDefaultSystemInstruction()
	}
	promptTemplate := template.Must(template.New("prompt").Parse(u.getDefaultPromptTemplate()))
	var promptBuffer bytes.Buffer
	err := promptTemplate.Execute(&promptBuffer, input.PullRequest)
	if err != nil {
		return "", err
	}
	res, err := u.sendPromptUseCase.Invoke(ctx, llmDomain.PromptInput{
		SystemInstruction: systemInstruction,
		Prompt:            promptBuffer.String(),
	})
	if err != nil {
		return "", err
	}
	return res, nil
}

func (u *reviewPullRequest) getDefaultSystemInstruction() string {
	return `
	Act as Software Enginner Expert with experience in multiple programming language.
	Always follow best practices by writing clean, modular code with proper security measures,
	and leveraging design patterns. Follow a software development principles: SOLID, DRY, KISS, YAGNI. Skip compliments.

	Input: Changes Patch. Changes patch changes between source branch and target branches in Pull request.
	Additional Context: PR Title, PR Description.

	Task: Review a file of source code, and the git diff of a set of changes made to that file on a Pull Request.
	- Do NOT provide general feedback, summaries, explanations of changes, or praises for making good additions. 
	- Focus solely on offering specific, objective insights based on the given context and refrain from making broad comments about potential impacts on the system or question intentions behind the changes.
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
) u.UseCase[domain.PullRequestReviewInput, string] {
	return &reviewPullRequest{
		sendPromptUseCase: sendPromptUseCase,
	}
}
