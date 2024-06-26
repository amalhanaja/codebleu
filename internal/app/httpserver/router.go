package httpserver

import (
	codeAssistantPresentation "codebleu/internal/presentation/handler/codeassistant"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(
	codeAssistantHandler codeAssistantPresentation.Handler,
) http.Handler {
	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("PONG"))
	})
	router.Get("/bitbucket/webhook", codeAssistantHandler.BitbucketWebhook)
	router.Post("/bitbucket/webhook", codeAssistantHandler.BitbucketWebhook)
	return router
}
