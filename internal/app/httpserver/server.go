package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, cfg AppConfig, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handler,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return server
}
