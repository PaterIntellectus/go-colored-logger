package logger

import "io"

type Options struct {
	Out    io.Writer
	Prefix string
	Flag   int
}
