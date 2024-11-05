// this package is responding for logging service information
package logger

import (
	"log"
	"os"

	"github.com/PaterIntellectus/almaray-gateway/internal/utils/ansi"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = getLogger(Options{
		Out:    getDefaultWriter(ansi.ForegroundBlue),
		Prefix: "INFO: ",
		Flag:   defaultWriterFlags,
	})

	WarningLogger = getLogger(Options{
		Out:    getDefaultWriter(ansi.ForegroundYellow),
		Prefix: "Warning: ",
		Flag:   defaultWriterFlags,
	})

	ErrorLogger = getLogger(Options{
		Out:    getDefaultWriter(ansi.ForegroundRed),
		Prefix: "Error: ",
		Flag:   defaultWriterFlags,
	})
}

func getLogger(options Options) *log.Logger {
	out := options.Out
	if out == nil {
		out = os.Stdout
	}

	return log.New(out, options.Prefix, options.Flag)
}

// Information logger interface
func Info(format string, v ...any) {
	InfoLogger.Printf(format, v...)
}

// Warning logger interface
func Warning(format string, v ...any) {
	WarningLogger.Printf(format, v...)
}

// Error logger interface
func Error(format string, v ...any) {
	ErrorLogger.Printf(format, v...)
}
