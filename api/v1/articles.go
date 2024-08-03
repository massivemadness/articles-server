package v1

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/massivemadness/articles-server/api/common"
	"github.com/massivemadness/articles-server/internal/articles"
	"net/http"
	"strconv"
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
		reqID := chi.URLParam(r, "id")
		artID, err := strconv.ParseInt(reqID, 10, 32)
		if err != nil {
			fmt.Printf("conversion error: %e\n", err)
		}
		fmt.Printf("artID = %d\n", artID)

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, common.HttpError{
			ErrorMessage: "Not implemented",
			ErrorCode:    common.ErrUnknown.Error(),
		})
	}
}
