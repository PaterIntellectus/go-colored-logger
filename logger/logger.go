package logger

type Logger interface {
	I(format string, v ...any)
	W(format string, v ...any)
	E(format string, v ...any)
	F(format string, v ...any)
}
