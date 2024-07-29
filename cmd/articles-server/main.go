package main

import (
	"fmt"
	"github.com/massivemadness/articles-server/api/v1"
	"github.com/massivemadness/articles-server/internal"
	"net/http"
	"time"
)

func main() {
	asv := &internal.ArticleServiceImpl{
		ServiceName: "Test 123",
	}

	httpServer := http.Server{
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", 8080),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      v1.NewRouter(asv),
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
