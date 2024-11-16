package logger

import (
	"io"
	"log"
	"os"

	"github.com/PaterIntellectus/go-colored-logger/ansi"
)

func New(options *options) Logger {
	var infoWriters []io.Writer
	var warningWriters []io.Writer
	var errorWriters []io.Writer
	var fatalWriters []io.Writer

	if options.Filename != "" {
		file, err := os.OpenFile(
			options.Filename,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0666,
		)

		if err != nil {
			log.Fatalf("logger: Cannot open file for logging - '%s'", options.Filename)
		}

		infoWriters = append(infoWriters, file)
		warningWriters = append(warningWriters, file)
		errorWriters = append(errorWriters, file)
		fatalWriters = append(fatalWriters, file)
	}

	if options.IsStdout {
		infoWriters = append(infoWriters, NewAnsiWriter(os.Stdout, ansi.ForegroundBrightBlue))
		warningWriters = append(warningWriters, NewAnsiWriter(os.Stdout, ansi.ForegroundBrightYellow))
		errorWriters = append(errorWriters, NewAnsiWriter(os.Stderr, ansi.ForegroundBrightRed))
		fatalWriters = append(fatalWriters, NewAnsiWriter(os.Stderr, ansi.ForegroundBrightMagenta))
	}

	return &LoggerImpl{
		infoLogger:    log.New(io.MultiWriter(infoWriters...), "[INFO]  ", options.Flag),
		warningLogger: log.New(io.MultiWriter(warningWriters...), "[WARN]  ", options.Flag),
		errorLogger:   log.New(io.MultiWriter(errorWriters...), "[ERROR] ", options.Flag),
		fatalLogger:   log.New(io.MultiWriter(fatalWriters...), "[FATAL] ", options.Flag),
	}
}
