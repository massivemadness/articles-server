package repository

import (
	"context"

	"github.com/massivemadness/articles-server/internal/entity"
	"github.com/massivemadness/articles-server/internal/storage"
)

type ArticleRepository interface {
	GetAll(ctx context.Context) ([]entity.Article, error)
	GetById(ctx context.Context, articleID int64) (*entity.Article, error)
	Create(ctx context.Context, article *entity.Article) (int64, error)
	Delete(ctx context.Context, articleID int64) error
}

type articleRepositoryImpl struct {
	db *storage.Storage
}

func NewArticleRepo(db *storage.Storage) ArticleRepository {
	return &articleRepositoryImpl{db: db}
}

func (r *articleRepositoryImpl) GetAll(ctx context.Context) ([]entity.Article, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, description FROM tbl_articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]entity.Article, 0)
	for rows.Next() {
		var article entity.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Desc)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleRepositoryImpl) GetById(ctx context.Context, articleID int64) (*entity.Article, error) {
	var article entity.Article
	err := r.db.QueryRow(
		ctx,
		"SELECT id, title, description FROM tbl_articles WHERE id = $1",
		articleID,
	).Scan(&article.ID, &article.Title, &article.Desc)

	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *articleRepositoryImpl) Create(ctx context.Context, article *entity.Article) (int64, error) {
	var articleID int64
	err := r.db.QueryRow(
		ctx,
		"INSERT INTO tbl_articles (title, description) VALUES ($1, $2) RETURNING id",
		article.Title,
		article.Desc,
	).Scan(&articleID)

	if err != nil {
		return 0, err
	}
	return articleID, nil
}

func (r *articleRepositoryImpl) Delete(ctx context.Context, articleID int64) error {
	_, err := r.db.Exec(ctx, "DELETE FROM tbl_articles WHERE id = $1", articleID)
	return err
}
