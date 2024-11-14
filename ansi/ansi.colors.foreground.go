package ansi

import "fmt"

type ForegroundColor Code

const (
	ForegroundBlack   ForegroundColor = 30
	ForegroundRed     ForegroundColor = 31
	ForegroundGreen   ForegroundColor = 32
	ForegroundYellow  ForegroundColor = 33
	ForegroundBlue    ForegroundColor = 34
	ForegroundMagenta ForegroundColor = 35
	ForegroundCyan    ForegroundColor = 36
	ForegroundWhite   ForegroundColor = 37

	ForegroundBrightBlack   ForegroundColor = 90
	ForegroundBrightRed     ForegroundColor = 91
	ForegroundBrightGreen   ForegroundColor = 92
	ForegroundBrightYellow  ForegroundColor = 93
	ForegroundBrightBlue    ForegroundColor = 94
	ForegroundBrightMagenta ForegroundColor = 95
	ForegroundBrightCyan    ForegroundColor = 96
	ForegroundBrightWhite   ForegroundColor = 97
)

func (fc ForegroundColor) String() string {
	return Code(fc).String()
}

func GetForegroundColor(r int, g int, b int) string {
	return fmt.Sprintf(
		"\x1b[38;2;%d;%d;%dm",
		r, g, b,
	)
}
