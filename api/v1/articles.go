package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/massivemadness/articles-server/api/common"
	"net/http"
	"strconv"
)

func GetArticlesHandler(wrapper *common.Wrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := wrapper.ArticleService.GetArticles()
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, common.HttpError{
				ErrorMessage: "Not found",
				ErrorCode:    common.ErrNotFound.Error(),
			})
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, ArticlesResponse{
			Articles: articles,
		})
	}
}

func GetArticleHandler(wrapper *common.Wrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := chi.URLParam(r, "id")
		artID, err := strconv.ParseInt(reqID, 10, 64)

		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, common.HttpError{
				ErrorMessage: "Invalid article ID",
				ErrorCode:    common.ErrInvalid.Error(),
			})
			return
		}

		article, err := wrapper.ArticleService.GetArticle(artID)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, common.HttpError{
				ErrorMessage: "Article not found",
				ErrorCode:    common.ErrInvalid.Error(),
			})
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, ArticleResponse{
			ID:          article.ID,
			Title:       article.Title,
			Description: article.Desc,
		})
	}
}
