package v1

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/massivemadness/articles-server/internal/articles"
	"net/http"
)

func NewRouter(asv articles.ArticleService) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/articles", GetArticlesHandler(asv))
		r.Get("/articles/{id}", GetArticleHandler(asv))
	})
	return r
}

func GetArticlesHandler(asv articles.ArticleService) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		fmt.Printf("Received v1/articles request: %v", asv.GetArticles())
	}
}

func GetArticleHandler(asv articles.ArticleService) func(http.ResponseWriter, *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		fmt.Printf("Received v1/articles/{id} request: %v", asv.GetArticle("123"))
	}
}
