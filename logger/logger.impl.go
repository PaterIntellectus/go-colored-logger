package logger

import (
	"log"
)

type LoggerImpl struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
}

func (logger *LoggerImpl) I(format string, v ...any) {
	logger.infoLogger.Printf(format, v...)
}

func (logger *LoggerImpl) W(format string, v ...any) {
	logger.warningLogger.Printf(format, v...)
}

func (logger *LoggerImpl) E(format string, v ...any) {
	logger.errorLogger.Printf(format, v...)
}

func (logger *LoggerImpl) F(format string, v ...any) {
	logger.fatalLogger.Fatalf(format, v...)
}
