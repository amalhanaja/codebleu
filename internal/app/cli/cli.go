package cli

import (
	"codebleu/internal/domain/codeassistant"
	"codebleu/internal/infrastructure/factories/gitrepo"
	"codebleu/internal/infrastructure/factories/llm"
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	codeAssistantUseCase "codebleu/internal/usecase/codeassistant"
	gitRepoUseCase "codebleu/internal/usecase/gitrepo"
	llmUseCase "codebleu/internal/usecase/llm"

	"github.com/urfave/cli/v2"
)

func NewCliApp() *cli.App {
	cliApp := cli.NewApp()
	cliApp.Name = "Codebleu"
	cliApp.Usage = "Review PR / MR Diff Changes"
	cliApp.Description = "Pull Request / Merge Request reviewer agent"
	cliApp.Version = "v0.0.1"
	cliApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "model",
			Usage:    `uses model to review pull request (options: "gemini-1.5-flash" (default), "gemini-1.5-pro", "gemini-1.0-pro")`,
			Aliases:  []string{"m"},
			EnvVars:  []string{"MODEL"},
			Value:    "gemini-1.5-flash",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "repository",
			Usage:    `hosted remote repository provider name (options: "bitbucket", "github")`,
			Aliases:  []string{"r"},
			EnvVars:  []string{"REPOSITORY_PROVIDER"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "id",
			Usage:    "pull request id",
			EnvVars:  []string{"PULL_REQUEST_ID"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "system-instruction",
			Usage:   "Custom system instruction for review pull request diff chages",
			EnvVars: []string{"SYSTEM_INSTRUCTION"},
		},
	}
	cliApp.Action = action

	return cliApp
}

func action(ctx *cli.Context) error {
	model := ctx.String("model")
	if model == "" {
		return errors.New("please provide model to review pull request")
	}
	remoteRepoProvider := ctx.String("repository")
	if remoteRepoProvider == "" {
		return errors.New("please provide repository provider (ex. github, bitbucket)")
	}
	remoteRepo, err := gitrepo.NewRepository(remoteRepoProvider)
	if err != nil {
		return err
	}
	llmRepo, err := llm.NewRepository(model)
	if err != nil {
		return err
	}
	systemInstructionPath := ctx.String("system-instruction")
	systemInstruction := ""
	if systemInstructionPath != "" {
		content, err := ioutil.ReadFile(systemInstructionPath)
		if err != nil {
			return errors.Join(err, fmt.Errorf("failed read file %s", systemInstructionPath))
		}
		systemInstruction = string(content)
	}
	getPullRequest := gitRepoUseCase.GetPullRequest(remoteRepo)
	postPullRequestComment := gitRepoUseCase.PostPullRequestComment(remoteRepo)
	sendPromptUseCase := llmUseCase.SendPromptUseCase(llmRepo)
	reviewPullRequest := codeAssistantUseCase.ReviewPullRequest(sendPromptUseCase)
	reviewAndCommentPullRequest := codeAssistantUseCase.ReviewAndCommentPullRequest(getPullRequest, reviewPullRequest, postPullRequestComment)
	if _, err := reviewAndCommentPullRequest.Invoke(
		context.Background(),
		codeassistant.ReviewAndCommentPullRequestInput{
			PullRequestId:     ctx.String("id"),
			SystemInstruction: systemInstruction,
		},
	); err != nil {
		return err
	}
	return nil
}
