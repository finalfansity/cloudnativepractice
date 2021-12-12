package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func LogInit(logLevel string) {
	logger = logrus.New()

	level := logrus.DebugLevel
	logger.SetLevel(logrus.ErrorLevel)
	switch {
	case logLevel == "debug":
		level = logrus.DebugLevel
	case logLevel == "info":
		level = logrus.InfoLevel
	case logLevel == "error":
		level = logrus.ErrorLevel
	default:
		level = logrus.DebugLevel
	}
	logger.SetOutput(os.Stdout)
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(level)

}

func Println(v ...interface{}) {
	logger.Info(v)
}

func Error(v ...interface{}) {
	logger.Error(v)
}

func Debug(v ...interface{}) {
	logger.Debug(v)
}
