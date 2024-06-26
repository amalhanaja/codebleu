package codeassistant

import (
	domain "codebleu/internal/domain/codeassistant"
	u "codebleu/internal/usecase"
	"context"
	"fmt"
)

type reviewPullRequest struct {
	sendPromptUseCase u.UseCase[string, string]
}

// Invoke implements usecase.UseCase.
func (u *reviewPullRequest) Invoke(ctx context.Context, input domain.PullRequestReviewInput) (string, error) {
	prompt := fmt.Sprintf(`
	Act as Principal Android Engineer Expert with experience in multiple programming language. 
	Always follow best practices by writing clean, modular code with proper security measures,
	and leveraging design patterns.

	General Functionality:\n
	- Does the code achieve the intended functionality as described in the pull request description?\b
    - Are there any edge cases that haven't been considered?\n
	- Is the new functionality well-integrated with the existing codebase?
	Code Quality:\n
	- Is the code well-formatted and easy to read? Does it follow the project's coding style guidelines?\n
	- Are there any opportunities for code simplification or refactoring?\n
	- Are comments included to explain complex logic?\n
	- Are meaningful variable, method, and class names used?\n
	- Can the code be easily tested?\n
	Performance:\n
	- Could the code be optimized for better performance?\n
	- Are there any potential memory leaks or resource leaks?\n
	Specific Code Suggestions:\n
	- If there are any code suggestions, mention specific code suggestion in detail.\n
	Potential issues:\n
	- Are there any potential backward compatibility issues?\n
	- Are there any potential security vulnerabilities introduced by the changes?\n
	
	Help me to review and analyze Pull request and Give constructive and positive feedback!.

	Here is Detail of pull request:

	Pull Request Title: %s\n
	Pull Request Description: \n%s\n

	Changes Patch:\n%s`, input.Title, input.Description, input.DiffPatch)
	res, err := u.sendPromptUseCase.Invoke(ctx, prompt)
	if err != nil {
		return "", err
	}
	return res, nil
}

func ReviewPullRequest(
	sendPromptUseCase u.UseCase[string, string],
) u.UseCase[domain.PullRequestReviewInput, string] {
	return &reviewPullRequest{
		sendPromptUseCase: sendPromptUseCase,
	}
}
