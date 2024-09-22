package articles

import (
	"context"

	"github.com/massivemadness/articles-server/internal/config"
	"github.com/massivemadness/articles-server/internal/entity"
	"github.com/massivemadness/articles-server/internal/repository"
	"go.uber.org/zap"
)

type ArticleService interface {
	GetArticles(ctx context.Context) ([]entity.Article, error)
	GetArticle(ctx context.Context, articleID int64) (*entity.Article, error)
	CreateArticle(ctx context.Context, article *entity.Article) (int64, error)
	DeleteArticle(ctx context.Context, articleID int64) error
}

type articleServiceImpl struct {
	repo   repository.ArticleRepository
	cfg    *config.Config
	logger *zap.Logger
}

func NewService(
	repository repository.ArticleRepository,
	cfg *config.Config,
	logger *zap.Logger,
) ArticleService {
	return &articleServiceImpl{
		repo:   repository,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *articleServiceImpl) GetArticles(ctx context.Context) ([]entity.Article, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *articleServiceImpl) GetArticle(ctx context.Context, articleID int64) (*entity.Article, error) {
	article, err := s.repo.GetById(ctx, articleID)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *articleServiceImpl) CreateArticle(ctx context.Context, article *entity.Article) (int64, error) {
	articleID, err := s.repo.Create(ctx, article)
	if err != nil {
		return 0, err
	}
	return articleID, nil
}

func (s *articleServiceImpl) DeleteArticle(ctx context.Context, articleID int64) error {
	return s.repo.Delete(ctx, articleID)
}
