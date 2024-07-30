package articles

import (
	"github.com/massivemadness/articles-server/internal/config"
	"go.uber.org/zap"
)

type ArticleService interface {
	GetArticles() string
	GetArticle(articleID string) string
}

type ArticleServiceImpl struct {
	cfg       *config.Config
	zapLogger *zap.Logger
}

func New(
	cfg *config.Config,
	zapLogger *zap.Logger,
) ArticleService {
	return &ArticleServiceImpl{
		cfg:       cfg,
		zapLogger: zapLogger,
	}
}

func (s *ArticleServiceImpl) GetArticles() string {
	return "[1, 2, 3, 4, 5]"
}

func (s *ArticleServiceImpl) GetArticle(articleID string) string {
	return articleID
}
