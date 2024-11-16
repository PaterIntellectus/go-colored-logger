package ansi

import "fmt"

const (
	ForegroundBlack   Code = 30
	ForegroundRed     Code = 31
	ForegroundGreen   Code = 32
	ForegroundYellow  Code = 33
	ForegroundBlue    Code = 34
	ForegroundMagenta Code = 35
	ForegroundCyan    Code = 36
	ForegroundWhite   Code = 37

	ForegroundBrightBlack   Code = 90
	ForegroundBrightRed     Code = 91
	ForegroundBrightGreen   Code = 92
	ForegroundBrightYellow  Code = 93
	ForegroundBrightBlue    Code = 94
	ForegroundBrightMagenta Code = 95
	ForegroundBrightCyan    Code = 96
	ForegroundBrightWhite   Code = 97
)

func ForegroundFromRGB(r int, g int, b int) string {
	return fmt.Sprintf(
		"\x1b[38;2;%d;%d;%dm",
		r, g, b,
	)
}
