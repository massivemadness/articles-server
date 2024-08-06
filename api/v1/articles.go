package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/massivemadness/articles-server/api/common"
	"github.com/massivemadness/articles-server/internal/articles"
	"net/http"
	"strconv"
)

func GetArticlesHandler(wrapper *common.Wrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbArticles, err := wrapper.ArticleService.GetArticles()
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
			Articles: dbArticles,
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

func CreateArticleHandler(wrapper *common.Wrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var createArticleRequest CreateArticleRequest

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&createArticleRequest)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, common.HttpError{
				ErrorMessage: "Invalid request body",
				ErrorCode:    common.ErrDecode.Error(),
			})
			return
		}
		
		// TODO validation

		article := articles.Article{
			ID:    0,
			Title: createArticleRequest.Title,
			Desc:  createArticleRequest.Description,
		}

		articleID, err := wrapper.ArticleService.CreateArticle(article)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, common.HttpError{
				ErrorMessage: "An error occurred",
				ErrorCode:    common.ErrUnknown.Error(),
			})
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, CreateArticleResponse{ID: articleID})
	}
}
