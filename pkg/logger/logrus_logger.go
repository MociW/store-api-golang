package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/MociW/store-api-golang/config"
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

func NewLogger(cfg *config.Config) *logrus.Logger {
	logger := logrus.New()

	level, exist := fieldsLogrusLevelMap[cfg.Logger.Level]
	if !exist {
		fmt.Printf("Invalid log level '%s', defaulting to 'debug'\n", cfg.Logger.Level)
		level = logrus.DebugLevel
	}

	if cfg.Logger.Encoding != "text" && cfg.Logger.Encoding != "json" {
		fmt.Printf("Invalid log encoding '%s', defaulting to 'json'\n", cfg.Logger.Encoding)
		cfg.Logger.Encoding = "json"
	}

	var formatter logrus.Formatter
	if cfg.Logger.Encoding == "text" {
		formatter = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			DisableColors:   false,
			// ForceColors:     true,
			CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
				file = fmt.Sprintf("%s:%d", f.File, f.Line)
				return function, filepath.Base(file)
			},
		}
	} else {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
			CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
				file = fmt.Sprintf("%s:%d", f.File, f.Line)
				return function, filepath.Base(file)
			},
		}
	}

	logger.SetLevel(level)
	logger.SetFormatter(formatter)
	logger.SetReportCaller(cfg.Logger.Caller)
	logger.SetOutput(os.Stdout)

	return logger
}
