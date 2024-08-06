package repository

import (
	"github.com/massivemadness/articles-server/internal/entity"
)

type ArticleRepository interface {
	GetArticles() ([]string, error)
	GetArticle(articleID int64) (entity.Article, error)
	CreateArticle(article entity.Article) (int64, error)
}

type articleRepositoryImpl struct {
	// TODO db sql.DB
}

func NewArticleRepo() ArticleRepository {
	return &articleRepositoryImpl{}
}

func (r *articleRepositoryImpl) GetArticles() ([]string, error) {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, nil
}

func (r *articleRepositoryImpl) GetArticle(articleID int64) (entity.Article, error) {
	return entity.Article{
		ID:    articleID,
		Title: "Lorem ipsum",
		Desc:  "Lorem ipsum dolor sit amet",
	}, nil
}

func (r *articleRepositoryImpl) CreateArticle(article entity.Article) (int64, error) {
	return article.ID, nil
}
