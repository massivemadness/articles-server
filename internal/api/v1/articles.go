package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/massivemadness/articles-server/internal/api/server"
	"github.com/massivemadness/articles-server/internal/entity"
	"net/http"
	"strconv"
)

func GetArticlesHandler(wrapper *server.Wrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbArticles, err := wrapper.ArticleService.GetArticles()
		if err != nil {
			server.ErrorJSON(w, r, http.StatusBadRequest, server.ErrNotFound)
			return
		}

		articles := make([]ArticleResponse, len(dbArticles))
		for i, article := range dbArticles {
			articles[i] = ArticleResponse{
				ID:          article.ID,
				Title:       article.Title,
				Description: article.Desc,
			}
		}

		response := ArticlesResponse{Articles: articles}
		server.ResponseJSON(w, r, http.StatusOK, response)
	}
}

func GetArticleHandler(wrapper *server.Wrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := chi.URLParam(r, "id")
		artID, err := strconv.ParseInt(reqID, 10, 64)
		if err != nil {
			server.ErrorJSON(w, r, http.StatusBadRequest, server.ErrDecode)
			return
		}

		article, err := wrapper.ArticleService.GetArticle(artID)
		if err != nil {
			server.ErrorJSON(w, r, http.StatusBadRequest, server.ErrNotFound)
			return
		}

		response := ArticleResponse{
			ID:          article.ID,
			Title:       article.Title,
			Description: article.Desc,
		}
		server.ResponseJSON(w, r, http.StatusOK, response)
	}
}

func CreateArticleHandler(wrapper *server.Wrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createArticleRequest CreateArticleRequest
		err := json.NewDecoder(r.Body).Decode(&createArticleRequest)
		if err != nil {
			server.ErrorJSON(w, r, http.StatusBadRequest, server.ErrDecode)
			return
		}

		err = wrapper.Validator.Struct(createArticleRequest)
		if err != nil {
			server.ErrorJSON(w, r, http.StatusBadRequest, server.ErrDecode)
			return
		}

		article := &entity.Article{
			ID:    0,
			Title: createArticleRequest.Title,
			Desc:  createArticleRequest.Description,
		}

		articleID, err := wrapper.ArticleService.CreateArticle(article)
		if err != nil {
			server.ErrorJSON(w, r, http.StatusBadRequest, server.ErrUnknown)
			return
		}

		response := CreateArticleResponse{ID: articleID}
		server.ResponseJSON(w, r, http.StatusCreated, response)
	}
}
