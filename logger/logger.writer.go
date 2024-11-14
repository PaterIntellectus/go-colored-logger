package logger

import (
	"io"

	"github.com/PaterIntellectus/go-colored-logger/ansi"
)

// don't use colored writer with files
type ColoredWriter struct {
	writer          io.Writer
	foregroundColor ansi.ForegroundColor
	backgroundColor ansi.BackgroundColor
}

func NewColoredWriter(
	writer io.Writer,
	foregroundColor ansi.ForegroundColor,
	backgroundColor ansi.BackgroundColor,
) *ColoredWriter {
	return &ColoredWriter{
		writer:          writer,
		foregroundColor: foregroundColor,
		backgroundColor: backgroundColor,
	}
}

func (cw *ColoredWriter) Write(p []byte) (n int, err error) {
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
