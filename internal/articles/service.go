package articles

import (
	"github.com/massivemadness/articles-server/internal/config"
	"github.com/massivemadness/articles-server/internal/entity"
	"github.com/massivemadness/articles-server/internal/repository"
	"go.uber.org/zap"
)

type ArticleService interface {
	GetArticles() ([]entity.Article, error)
	GetArticle(articleID int64) (entity.Article, error)
	CreateArticle(article entity.Article) (int64, error)
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

func (s *articleServiceImpl) GetArticles() ([]entity.Article, error) {
	data, err := s.repo.GetArticles()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *articleServiceImpl) GetArticle(articleID int64) (entity.Article, error) {
	return entity.Article{
		ID:    articleID,
		Title: "Lorem ipsum",
		Desc:  "Lorem ipsum dolor sit amet",
	}, nil
}

func (s *articleServiceImpl) CreateArticle(_ entity.Article) (int64, error) {
	return 0, nil
}
