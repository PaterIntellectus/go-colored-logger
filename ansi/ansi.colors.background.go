package ansi

import "fmt"

const (
	BackgroundBlack   Code = 40
	BackgroundRed     Code = 41
	BackgroundGreen   Code = 42
	BackgroundYellow  Code = 43
	BackgroundBlue    Code = 44
	BackgroundMagenta Code = 45
	BackgroundCyan    Code = 46
	BackgroundWhite   Code = 47

	BackgroundBrightBlack   Code = 100
	BackgroundBrightRed     Code = 101
	BackgroundBrightGreen   Code = 102
	BackgroundBrightYellow  Code = 103
	BackgroundBrightBlue    Code = 104
	BackgroundBrightMagenta Code = 105
	BackgroundBrightCyan    Code = 106
	BackgroundBrightWhite   Code = 107
)

func BackgroundFromRGB(r int, g int, b int) string {
	return fmt.Sprintf(
		"\x1b[48;2;%d;%d;%dm",
		r, g, b,
	)
}
