package internal

type ArticleService interface {
	GetArticles() string
	GetArticle(articleID string) string
}

type ArticleServiceImpl struct {
	ServiceName string
}

func (s *ArticleServiceImpl) GetArticles() string {
	return "[1, 2, 3, 4, 5]"
}

func (s *ArticleServiceImpl) GetArticle(articleID string) string {
	return articleID
}
