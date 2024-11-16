package logger

type options struct {
	Filename string
	IsStdout bool
	Flag     int
}

func NewOptions(
	filename string,
	isStdout bool,
	flag int,
) *options {
	return &options{
		Filename: filename,
		IsStdout: isStdout,
		Flag:     flag,
	}
}

func NewOptionsWithDefaultFlags(
	filename string,
	isStdout bool,
) *options {
	return &options{
		Filename: filename,
		IsStdout: isStdout,
		Flag:     DefaultFlags,
	}
}
