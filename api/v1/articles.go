package v1

import (
	"github.com/go-chi/render"
	"github.com/massivemadness/articles-server/api/common"
	"github.com/massivemadness/articles-server/internal/articles"
	"net/http"
)

func GetArticlesHandler(asv articles.ArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, ArticlesResponse{
			Articles: asv.GetArticles(),
		})
	}
}

func GetArticleHandler(_ articles.ArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, common.HttpError{
			ErrorMessage: "Not implemented",
			ErrorCode:    common.ErrUnknown.Error(),
		})
	}
}
