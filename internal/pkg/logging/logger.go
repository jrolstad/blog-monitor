package logging

import (
	"go.uber.org/zap"
	"sync"
)

func LogEvent(message string, keysAndValues ...interface{}) {
	logger := getLogger()
	logger.Infow(message, addLogType("event", keysAndValues)...)
}

func LogDependency(message string, keysAndValues ...interface{}) {
	logger := getLogger()
	logger.Infow(message, addLogType("dependency", keysAndValues)...)
}

func LogError(err error, keysAndValues ...interface{}) {
	logger := getLogger()
	logger.Errorw(err.Error(), addLogType("error", keysAndValues)...)
}

func LogPanic(err error, keysAndValues ...interface{}) {
	logger := getLogger()
	logger.Panicw(err.Error(), addLogType("error", keysAndValues)...)
}

var logger *zap.SugaredLogger
var locker = &sync.Mutex{}

func getLogger() *zap.SugaredLogger {
	if logger == nil {
		locker.Lock()

		prodLogger, _ := zap.NewProduction()
		logger = prodLogger.Sugar()

		locker.Unlock()
	}

	return logger
}

func addLogType(logType string, keysAndValues []interface{}) []interface{} {
	return append(keysAndValues, "logType", logType)
}
