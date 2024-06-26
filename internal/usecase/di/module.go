package di

import (
	"codebleu/internal/usecase/codeassistant"
	"codebleu/internal/usecase/gitrepo"
	"codebleu/internal/usecase/llm"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"usecases",
		fx.Provide(
			fx.Annotate(
				gitrepo.GetPullRequest,
				fx.ResultTags(`name:"GetPullRequest"`),
			),
			fx.Annotate(
				gitrepo.PostPullRequestComment,
				fx.ResultTags(`name:"PostPullRequestComment"`),
			),
			fx.Annotate(
				llm.SendPromptUseCase,
				fx.ResultTags(`name:"SendPromptUseCase"`),
			),
			fx.Annotate(
				codeassistant.ReviewPullRequest,
				fx.ParamTags(`name:"SendPromptUseCase"`),
				fx.ResultTags(`name:"ReviewPullRequest"`),
			),
			fx.Annotate(
				codeassistant.ReviewAndCommentPullRequest,
				fx.ParamTags(
					`name:"GetPullRequest"`,
					`name:"ReviewPullRequest"`,
					`name:"PostPullRequestComment"`,
				),
				fx.ResultTags(`name:"ReviewAndCommentPullRequest"`),
			),
		),
	)
}
