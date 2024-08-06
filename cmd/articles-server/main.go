package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/massivemadness/articles-server/api"
	"github.com/massivemadness/articles-server/api/server"
	"github.com/massivemadness/articles-server/internal/articles"
	"github.com/massivemadness/articles-server/internal/config"
	"github.com/massivemadness/articles-server/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	validate := validator.New()
	zapLogger := logger.NewLogger(cfg.Env)
	articleService := articles.NewService(cfg, zapLogger)

	wrapper := &server.Wrapper{
		ArticleService: articleService,
		Cfg:            cfg,
		Validator:      validate,
		Logger:         zapLogger,
	}

	httpServer := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HttpServer.Address, cfg.HttpServer.Port),
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
		Handler:      api.NewRouter(wrapper),
	}

	zapLogger.Info("Starting http server")

	go func() {
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			zapLogger.Error("HTTP server error", zap.Error(err))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		zapLogger.Error("HTTP shutdown error", zap.Error(err))
	}

	zapLogger.Info("Server stopped")
}
