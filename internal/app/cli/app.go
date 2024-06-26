package cli

import (
	infraBitbucket "codebleu/internal/infrastructure/http/bitbucket"
	infraGemini "codebleu/internal/infrastructure/thirdparty/gemini"
	codeAssistantUseCase "codebleu/internal/usecase/codeassistant"
	gitRepoUseCase "codebleu/internal/usecase/gitrepo"
	llmUseCase "codebleu/internal/usecase/llm"
	"context"
)

func Run(pullRequestId string) {
	cfg := NewConfigFromEnv()
	bitbucketClient := infraBitbucket.NewClient(
		cfg.BitbucketRepositoryAccessConfig.Workspace,
		cfg.BitbucketRepositoryAccessConfig.RepoSlug,
		cfg.BitbucketRepositoryAccessConfig.AccessToken,
	)
	geminiClient := infraGemini.NewClient(cfg.GeminiConfig.ApiKey)
	getPullRequest := gitRepoUseCase.GetPullRequest(bitbucketClient)
	postPullRequestComment := gitRepoUseCase.PostPullRequestComment(bitbucketClient)
	sendPromptUseCase := llmUseCase.SendPromptUseCase(geminiClient)
	reviewPullRequest := codeAssistantUseCase.ReviewPullRequest(sendPromptUseCase)
	reviewAndCommentPullRequest := codeAssistantUseCase.ReviewAndCommentPullRequest(getPullRequest, reviewPullRequest, postPullRequestComment)

	_, err := reviewAndCommentPullRequest.Invoke(context.Background(), pullRequestId)
	if err != nil {
		panic(err)
	}
}
