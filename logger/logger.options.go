package logger

// To disable file logging set filename to ""
type Options struct {
	Filename string
	Stdout   bool
}

func NewOptions(
	filename string,
	stdout bool,
) *Options {
	return &Options{
		Filename: filename,
		Stdout:   stdout,
	}
}
