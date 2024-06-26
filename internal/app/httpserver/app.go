package httpserver

import (
	gitRepoDomain "codebleu/internal/domain/gitrepo"
	llmDomain "codebleu/internal/domain/llm"
	infraBitbucket "codebleu/internal/infrastructure/http/bitbucket"
	infraGemini "codebleu/internal/infrastructure/thirdparty/gemini"
	codeAssistantPresentation "codebleu/internal/presentation/handler/codeassistant"
	useCaseInjection "codebleu/internal/usecase/di"
	"net/http"

	"github.com/go-playground/webhooks/v6/bitbucket"
	"go.uber.org/fx"
)

// Injections
var (
	configModule = fx.Module("config", fx.Provide(NewConfigFromEnv))

	domainModule = fx.Module(
		"domain",
		fx.Provide(func(cfg AppConfig) gitRepoDomain.Repository {
			return infraBitbucket.NewClient(
				cfg.BitbucketRepositoryAccessConfig.Workspace,
				cfg.BitbucketRepositoryAccessConfig.RepoSlug,
				cfg.BitbucketRepositoryAccessConfig.AccessToken,
			)
		}),
		fx.Provide(func(cfg AppConfig) llmDomain.Repository {
			return infraGemini.NewClient(cfg.GeminiConfig.ApiKey)
		}),
	)

	thirdPartyModule = fx.Module(
		"thirdParty",
		fx.Provide(func(cfg AppConfig) (*bitbucket.Webhook, error) {
			return bitbucket.New(bitbucket.Options.UUID(cfg.BitbucketWebhookConfig.Secret))
		}),
	)

	presentationModule = fx.Module(
		"presentation",
		fx.Provide(
			fx.Annotate(
				codeAssistantPresentation.NewHandler,
				fx.ParamTags(`name:"ReviewAndCommentPullRequest"`),
			),
		),
	)

	appModule = fx.Module("application", fx.Provide(NewHttpServer, NewRouter))
)

func Run() {

	appStartInvoker := fx.Invoke(func(*http.Server) {}) // Invoke to start application
	fx.New(
		configModule,
		domainModule,
		thirdPartyModule,
		useCaseInjection.Module(),
		presentationModule,
		appModule,
		appStartInvoker,
	).Run()
}
