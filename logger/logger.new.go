package logger

import (
	"io"
	"log"
	"os"

	"github.com/PaterIntellectus/go-colored-logger/ansi"
)

func New(options *Options) *LoggerImpl {
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

	if options.Stdout {
		infoWriters = append(infoWriters, NewColoredWriter(
			os.Stdout,
			ansi.ForegroundBrightBlue,
			ansi.BackgroundColor(ansi.ResetAll),
		))
		warningWriters = append(warningWriters, NewColoredWriter(
			os.Stdout,
			ansi.ForegroundBrightYellow,
			ansi.BackgroundColor(ansi.ResetAll),
		))
		errorWriters = append(errorWriters, NewColoredWriter(
			os.Stdout,
			ansi.ForegroundBrightRed,
			ansi.BackgroundColor(ansi.ResetAll),
		))
		fatalWriters = append(fatalWriters, NewColoredWriter(
			os.Stdout,
			ansi.ForegroundMagenta,
			ansi.BackgroundColor(ansi.ResetAll),
		))
	}

	defaultWriterFlags := log.Ldate | log.Ltime

	return &LoggerImpl{
		infoLogger:    log.New(io.MultiWriter(infoWriters...), "[INFO] ", defaultWriterFlags),
		warningLogger: log.New(io.MultiWriter(warningWriters...), "[WARNING] ", defaultWriterFlags),
		errorLogger:   log.New(io.MultiWriter(errorWriters...), "[ERROR] ", defaultWriterFlags),
		fatalLogger:   log.New(io.MultiWriter(fatalWriters...), "[FATAL] ", defaultWriterFlags),
	}
}
