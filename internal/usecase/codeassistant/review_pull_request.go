package codeassistant

import (
	"bytes"
	domain "codebleu/internal/domain/codeassistant"
	u "codebleu/internal/usecase"
	"context"
	"text/template"
)

type reviewPullRequest struct {
	sendPromptUseCase u.UseCase[string, string]
}

// Invoke implements usecase.UseCase.
func (u *reviewPullRequest) Invoke(ctx context.Context, input domain.PullRequestReviewInput) (string, error) {
	promptTemplateString := input.ReviewPrompt
	if promptTemplateString == "" {
		promptTemplateString = u.getDefaultPromptTemplate()
	}
	promptTemplate := template.Must(template.New("prompt").Parse(promptTemplateString))
	var promptBuffer bytes.Buffer
	promptTemplate.Execute(&promptBuffer, input.PullRequest)
	res, err := u.sendPromptUseCase.Invoke(ctx, promptBuffer.String())
	if err != nil {
		return "", err
	}
	return res, nil
}

func (u *reviewPullRequest) getDefaultPromptTemplate() string {
	return `
	## PR Title:
	{{.Title}}
	## Description:
	{{.Description}}
	## Instructions:
	Act as Software Enginner Expert with experience in multiple programming language. 
	Always follow best practices by writing clean, modular code with proper security measures,
	and leveraging design patterns.

	Input: Changes Patch. Changes patch changes between source branch and target branches in Pull request.
	Additional Context: PR title, description.
	Task: Review changes patch for substantive issues using provided context and respond with comments if necessary.
	Output: Review comments in markdown format.
	- Use fenced code blocks using the relevant language identifier where applicable.
	- Don't annotate code snippets with line numbers. Format and indent code correctly.

	## Changes Patch:
	{{.DiffPatch}}
	`
}

func ReviewPullRequest(
	sendPromptUseCase u.UseCase[string, string],
) u.UseCase[domain.PullRequestReviewInput, string] {
	return &reviewPullRequest{
		sendPromptUseCase: sendPromptUseCase,
	}
}
