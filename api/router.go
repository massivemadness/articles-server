package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/massivemadness/articles-server/api/v1"
	"github.com/massivemadness/articles-server/internal/articles"
)

func NewRouter(asv articles.ArticleService) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/articles", v1.GetArticlesHandler(asv))
		r.Get("/articles/{id}", v1.GetArticleHandler(asv))
		// TODO create (POST)
		// TODO update (PATCH)
		// TODO delete (DELETE)
	})
	return r
}