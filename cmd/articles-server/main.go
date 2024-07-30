package main

import (
	"fmt"
	"github.com/massivemadness/articles-server/api/v1"
	"github.com/massivemadness/articles-server/internal/articles"
	"github.com/massivemadness/articles-server/internal/config"
	"net/http"
)

func main() {
	cfg := config.Load()

	// TODO create deps

	asv := articles.New(cfg)

	httpServer := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HttpServer.Address, cfg.HttpServer.Port),
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
		Handler:      v1.NewRouter(asv),
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
	}
}
