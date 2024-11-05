package logger

import (
	"io"

	"github.com/PaterIntellectus/almaray-gateway/internal/utils/ansi"
)

// don't use colored writer with files
type Writer struct {
	writer          io.Writer
	foregroundColor ansi.ForegroundColor
	backgroundColor ansi.BackgroundColor
}

func NewWriter(
	writer io.Writer,
	foregroundColor ansi.ForegroundColor,
	backgroundColor ansi.BackgroundColor,
) *Writer {
	return &Writer{
		writer:          writer,
		foregroundColor: foregroundColor,
		backgroundColor: backgroundColor,
	}
}

func (cw *Writer) Write(p []byte) (n int, err error) {
	if cw.foregroundColor == 0 {
		return cw.writer.Write(p)
	} else {
		return cw.writer.Write(
			[]byte(ansi.Apply(
				string(p),
				ansi.Code(cw.foregroundColor),
				ansi.Code(cw.backgroundColor),
			)),
		)
	}
}
