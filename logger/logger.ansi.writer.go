package logger

import (
	"io"

	"github.com/PaterIntellectus/go-colored-logger/ansi"
)

// Shouldn't be used with outs
// that doesn't recognize ansi codes
type AnsiWriter struct {
	writer io.Writer
	codes  []ansi.Code
}

func NewAnsiWriter(
	writer io.Writer,
	codes ...ansi.Code,
) *AnsiWriter {
	return &AnsiWriter{
		writer: writer,
		codes:  codes,
	}
}

func (aw *AnsiWriter) Write(p []byte) (n int, err error) {
	if len(aw.codes) > 0 {
		return aw.writer.Write([]byte(ansi.Apply(string(p), aw.codes...)))
	} else {
		return aw.writer.Write(p)
	}
}
