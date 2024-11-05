package logger

import (
	"io"
	"log"
	"os"

	"github.com/PaterIntellectus/almaray-gateway/internal/utils/ansi"
)

const (
	defaultWriterFlags = log.Ldate | log.Ltime
	defaultFilename    = "almaray-gateway.log"
)

var defaultFile *os.File

func getDefaultWriter(foregroundColor ansi.ForegroundColor) io.Writer {
	if defaultFile == nil {
		file, err := os.OpenFile(
			defaultFilename,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0666,
		)

		if err != nil {
			log.Fatalf("logger: Cannot open default file for logging - '%s'", defaultFilename)
		}

		defaultFile = file
	}

	return io.MultiWriter(
		defaultFile,
		NewWriter(
			os.Stdout,
			foregroundColor,
			ansi.BackgroundBlack,
		),
	)
}

// func init() {
// 	Init(Options{
// 		InfoLoggerOptions: Option{
// 			Out:    getDefaultWriter(ansi.ForegroundBlue),
// 			Prefix: "INFO: ",
// 			Flag:   defaultWriterFlags,
// 		},
// 		WarningLoggerOptions: Option{
// 			Out:    getDefaultWriter(ansi.ForegroundYellow),
// 			Prefix: "Warning: ",
// 			Flag:   defaultWriterFlags,
// 		},
// 		ErrorLoggerOptions: Option{
// 			Out:    getDefaultWriter(ansi.ForegroundRed),
// 			Prefix: "Error: ",
// 			Flag:   defaultWriterFlags,
// 		},
// 	})
// }
