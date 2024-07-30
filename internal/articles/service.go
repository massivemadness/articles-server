package articles

import "github.com/massivemadness/articles-server/internal/config"

type ArticleService interface {
	GetArticles() string
	GetArticle(articleID string) string
}

type ArticleServiceImpl struct {
	cfg *config.Config
}

func New(cfg *config.Config) ArticleService {
	return &ArticleServiceImpl{
		cfg: cfg,
	}
}

func (s *ArticleServiceImpl) GetArticles() string {
	return "[1, 2, 3, 4, 5]"
}

func (s *ArticleServiceImpl) GetArticle(articleID string) string {
	return articleID
}
