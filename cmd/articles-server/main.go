package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/massivemadness/articles-server/internal/api"
	"github.com/massivemadness/articles-server/internal/api/server"
	"github.com/massivemadness/articles-server/internal/articles"
	"github.com/massivemadness/articles-server/internal/config"
	"github.com/massivemadness/articles-server/internal/logger"
	"github.com/massivemadness/articles-server/internal/repository"
	"github.com/massivemadness/articles-server/internal/storage"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()
	validate := validator.New()
	zapLogger := logger.NewLogger(cfg.Env)
	db, err := storage.New(cfg)
	if err != nil {
		zapLogger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	articleRepository := repository.NewArticleRepo(db)
	articleService := articles.NewService(articleRepository, cfg, zapLogger)

	wrapper := &server.Wrapper{
		ArticleService: articleService,
		Cfg:            cfg,
		Validator:      validate,
		Logger:         zapLogger,
	}

	httpServer := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HttpServer.Address, cfg.HttpServer.PublicPort),
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
		Handler:      api.PublicRouter(wrapper),
	}

	zapLogger.Info("Starting http server...")

	go func() {
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			zapLogger.Fatal("HTTP server error", zap.Error(err))
		}
	}()

	httpPrivateServer := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HttpServer.Address, cfg.HttpServer.PrivatePort),
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
		Handler:      api.PrivateRouter(),
	}

	go func() {
		if err := httpPrivateServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			zapLogger.Fatal("HTTP server error private", zap.Error(err))
		}
	}()

	zapLogger.Info("Service is ready")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), cfg.HttpServer.ShutdownTimeout)
	defer shutdownRelease()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		zapLogger.Error("HTTP shutdown error", zap.Error(err))
	}

	if err := httpPrivateServer.Shutdown(shutdownCtx); err != nil {
		zapLogger.Error("HTTP shutdown error private", zap.Error(err))
	}

	zapLogger.Info("Server stopped")
}
