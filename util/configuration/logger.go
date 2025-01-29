package configuration

import (
	"go.uber.org/zap"
)

func Logger () *zap.Logger{
	lc := zap.NewDevelopmentConfig()
	lc.DisableStacktrace = true
	logger, _ := lc.Build()
	return logger
}