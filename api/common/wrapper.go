package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/massivemadness/articles-server/internal/articles"
	"github.com/massivemadness/articles-server/internal/config"
	"go.uber.org/zap"
)

type Wrapper struct {
	ArticleService articles.ArticleService
	Cfg            *config.Config
	Validator      *validator.Validate
	Logger         *zap.Logger
}
