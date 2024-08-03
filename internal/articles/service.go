package articles

import (
	"github.com/massivemadness/articles-server/internal/config"
	"go.uber.org/zap"
)

type ArticleService interface {
	GetArticles() ([]string, error)
	GetArticle(articleID int64) (Article, error)
}

type ArticleServiceImpl struct {
	cfg    *config.Config
	logger *zap.Logger
}

func NewService(
	cfg *config.Config,
	logger *zap.Logger,
) ArticleService {
	return &ArticleServiceImpl{
		cfg:    cfg,
		logger: logger,
	}
}

type Article struct {
	ID    int64
	Title string
	Desc  string
}

func (s *ArticleServiceImpl) GetArticles() ([]string, error) {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, nil
}

func (s *ArticleServiceImpl) GetArticle(articleID int64) (Article, error) {
	return Article{
		ID:    articleID,
		Title: "Lorem ipsum",
		Desc:  "Lorem ipsum dolor sit amet",
	}, nil
}
