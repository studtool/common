package logs

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: func() *logrus.Logger {
			log := logrus.StandardLogger()
			log.SetFormatter(&logrus.JSONFormatter{})
			return log
		}(),
	}
}

type LogFields struct {
	Component string
	Function  string
}

func (log *Logger) Debug(f *LogFields, args ...interface{}) {
	log.logger.Debug(args...)
}

func (log *Logger) Info(f *LogFields, args ...interface{}) {
	log.logger.Info(args...)
}

func (log *Logger) Warning(f *LogFields, args ...interface{}) {
	log.logger.Warn(args...)
}

func (log *Logger) Error(f *LogFields, args ...interface{}) {
	log.logger.Error(args...)
}

func (log *Logger) Fatal(f *LogFields, args ...interface{}) {
	log.logger.Fatal(args...)
}

func (log *Logger) mapFields(f *LogFields) logrus.Fields {
	return logrus.Fields{
		"component": f.Component,
		"function":  f.Function,
	}
}
