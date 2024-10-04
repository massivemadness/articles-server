package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	mw "github.com/massivemadness/articles-server/internal/api/middleware"
	"github.com/massivemadness/articles-server/internal/api/server"
	v1 "github.com/massivemadness/articles-server/internal/api/v1"
)

func PublicRouter(wrapper *server.Wrapper) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(mw.Logger(wrapper.Logger))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/articles", v1.GetArticlesHandler(wrapper))
		r.Get("/articles/{id}", v1.GetArticleHandler(wrapper))
		r.Post("/articles/create", v1.CreateArticleHandler(wrapper))
		r.Delete("/articles/delete/{id}", v1.DeleteArticleHandler(wrapper))
	})
	return r
}

func PrivateRouter() chi.Router {
	r := chi.NewRouter()

	r.Mount("/debug", middleware.Profiler())

	return r
}
