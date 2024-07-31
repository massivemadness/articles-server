package main

import (
	"fmt"
	"github.com/massivemadness/articles-server/api"
	"github.com/massivemadness/articles-server/internal/articles"
	"github.com/massivemadness/articles-server/internal/config"
	"github.com/massivemadness/articles-server/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	zapLogger := logger.NewLogger(cfg.Env)

	// TODO create deps

	asv := articles.New(cfg, zapLogger)

	httpServer := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HttpServer.Address, cfg.HttpServer.Port),
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
		Handler:      api.NewRouter(asv),
	}

	zapLogger.Info("Starting http server")

	serverError := httpServer.ListenAndServe()
	if serverError != nil {
		zapLogger.Error("Cannot start http server", zap.Error(serverError))
	}
}
