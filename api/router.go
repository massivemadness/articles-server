package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/massivemadness/articles-server/api/common"
	mw "github.com/massivemadness/articles-server/api/middleware"
	"github.com/massivemadness/articles-server/api/v1"
)

func NewRouter(wrapper *common.Wrapper) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(mw.Logger(wrapper.Logger))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/articles", v1.GetArticlesHandler(wrapper))
		r.Get("/articles/{id}", v1.GetArticleHandler(wrapper))
		// TODO create (POST)
		// TODO update (PATCH)
		// TODO delete (DELETE)
	})
	return r
}
