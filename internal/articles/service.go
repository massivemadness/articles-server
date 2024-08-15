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
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *articleServiceImpl) GetArticle(articleID int64) (entity.Article, error) {
	article, err := s.repo.GetById(articleID)
	if err != nil {
		return entity.Article{}, err
	}
	return article, nil
}

func (s *articleServiceImpl) CreateArticle(article entity.Article) (int64, error) {
	articleID, err := s.repo.Create(article)
	if err != nil {
		return 0, err
	}
	return articleID, nil
}
