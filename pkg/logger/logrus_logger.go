package logger

import (
	"github.com/MociW/store-api-golang/pkg/config"
	"github.com/sirupsen/logrus"
)

var fieldsLogrusLevelMap = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

func NewLogger(cfg config.Config) *logrus.Logger {
	logger := logrus.New()

	// level, exist := fieldsLogrusLevelMap[cfg.Logger.Level]
	// if !exist {
	// 	level = logrus.DebugLevel
	// }

	return logger
}
