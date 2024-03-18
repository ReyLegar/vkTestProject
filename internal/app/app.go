package app

import (
	"context"
	"net/http"
	"time"

	"github.com/ReyLegar/vkTestProject/internal/config"
)

type App struct {
	cfg *config.Config

	httpServer *http.Server
	mux        *http.ServeMux
}

func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	app := &App{
		cfg:        cfg,
		mux:        mux,
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Run(handler http.Handler) error {

	a.httpServer = &http.Server{
		Addr:           a.cfg.BindAddr,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return a.httpServer.ListenAndServe()
}
