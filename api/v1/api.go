package v1

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/massivemadness/articles-server/internal"
	"net/http"
)

func NewRouter(asv internal.ArticleService) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/articles", func(resp http.ResponseWriter, req *http.Request) {
			fmt.Printf("Received v1/articles request: %v", asv.GetArticles())
		})
		r.Get("/articles/{id}", func(resp http.ResponseWriter, req *http.Request) {
			fmt.Printf("Received v1/articles/{id} request: %v", asv.GetArticle("123"))
		})
	})
	return r
}
