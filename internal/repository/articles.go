package repository

import (
	"context"
	"github.com/massivemadness/articles-server/internal/entity"
	"github.com/massivemadness/articles-server/internal/storage"
)

type ArticleRepository interface {
	GetArticles() ([]entity.Article, error)
	GetArticle(articleID int64) (entity.Article, error)
	CreateArticle(article entity.Article) (int64, error)
}

type articleRepositoryImpl struct {
	db *storage.Storage
}

func NewArticleRepo(db *storage.Storage) ArticleRepository {
	return &articleRepositoryImpl{db: db}
}

func (r *articleRepositoryImpl) GetArticles() ([]entity.Article, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM tbl_articles;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]entity.Article, 0)
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		article := entity.Article{
			ID:    values[0].(int64),
			Title: values[1].(string),
			Desc:  values[2].(string),
		}
		articles = append(articles, article)
	}

	return articles, nil
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
