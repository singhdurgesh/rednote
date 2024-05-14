package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	ZapLogger *zap.Logger
	ZapSugar  *zap.SugaredLogger

	LogrusLogger *logrus.Logger
)

func Init() *logrus.Logger {
	LogrusLogger = InitLogrusLogger()

	// ZapLogger = InitZapLogger()
	// ZapSugar = ZapLogger.Sugar()
	return LogrusLogger
}
