package logger

import (
	"github.com/massivemadness/articles-server/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(env string) *zap.Logger {
	var logger *zap.Logger

	switch env {
	case config.EnvProd:
		logger = zap.Must(zap.NewProduction())
	default:
		zapConfig := zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger = zap.Must(zapConfig.Build())
	}

	return logger
}
