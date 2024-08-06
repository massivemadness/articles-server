package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	mw "github.com/massivemadness/articles-server/api/middleware"
	"github.com/massivemadness/articles-server/api/server"
	"github.com/massivemadness/articles-server/api/v1"
)

func NewRouter(wrapper *server.Wrapper) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(mw.Logger(wrapper.Logger))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/articles", v1.GetArticlesHandler(wrapper))
		r.Get("/articles/{id}", v1.GetArticleHandler(wrapper))
		r.Post("/articles/create", v1.CreateArticleHandler(wrapper))
		// TODO update (PATCH)
		// TODO delete (DELETE)
	})
	return r
}
