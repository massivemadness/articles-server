package v1

import (
	"github.com/go-chi/render"
	"github.com/massivemadness/articles-server/api"
	"github.com/massivemadness/articles-server/internal/articles"
	"net/http"
)

func GetArticlesHandler(asv articles.ArticleService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, ArticlesResponse{
			Articles: asv.GetArticles(),
		})
	}
}

func GetArticleHandler(_ articles.ArticleService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, api.HttpError{
			ErrorMessage: "Not implemented",
			ErrorCode:    api.ErrUnknown.Error(),
		})
	}
}
