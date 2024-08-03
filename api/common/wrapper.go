package common

import (
	"github.com/massivemadness/articles-server/internal/articles"
	"github.com/massivemadness/articles-server/internal/config"
	"go.uber.org/zap"
)

type Wrapper struct {
	ArticleService articles.ArticleService
	Cfg            *config.Config
	Logger         *zap.Logger
}
