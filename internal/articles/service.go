package articles

import (
	"github.com/massivemadness/articles-server/internal/config"
	"go.uber.org/zap"
)

type ArticleService interface {
	GetArticles() []string
	GetArticle(articleID string) string
}

type ArticleServiceImpl struct {
	cfg    *config.Config
	logger *zap.Logger
}

func New(
	cfg *config.Config,
	logger *zap.Logger,
) ArticleService {
	return &ArticleServiceImpl{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *ArticleServiceImpl) GetArticles() []string {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

func (s *ArticleServiceImpl) GetArticle(articleID string) string {
	return articleID
}
