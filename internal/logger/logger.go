package logger

import (
	"github.com/massivemadness/articles-server/internal/config"
	"go.uber.org/zap"
)

func NewZapLogger(env string) *zap.Logger {
	var logger *zap.Logger

	switch env {
	case config.EnvLocal:
		logger = zap.Must(zap.NewDevelopment())
	case config.EnvDev:
		logger = zap.Must(zap.NewDevelopment())
	case config.EnvProd:
		logger = zap.Must(zap.NewProduction())
	default:
		logger = zap.NewNop()
	}

	return logger
}
