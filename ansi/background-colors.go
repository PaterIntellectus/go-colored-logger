package ansi

import "fmt"

type BackgroundColor Code

const (
	BackgroundBlack   BackgroundColor = 40
	BackgroundRed     BackgroundColor = 41
	BackgroundGreen   BackgroundColor = 42
	BackgroundYellow  BackgroundColor = 43
	BackgroundBlue    BackgroundColor = 44
	BackgroundMagenta BackgroundColor = 45
	BackgroundCyan    BackgroundColor = 46
	BackgroundWhite   BackgroundColor = 47

	BackgroundBrightBlack   BackgroundColor = 100
	BackgroundBrightRed     BackgroundColor = 101
	BackgroundBrightGreen   BackgroundColor = 102
	BackgroundBrightYellow  BackgroundColor = 103
	BackgroundBrightBlue    BackgroundColor = 104
	BackgroundBrightMagenta BackgroundColor = 105
	BackgroundBrightCyan    BackgroundColor = 106
	BackgroundBrightWhite   BackgroundColor = 107
)

func (fc BackgroundColor) String() string {
	return Code(fc).String()
}

func GetBackgroundColor(r int, g int, b int) string {
	return fmt.Sprintf(
		"\x1b[48;2;%d;%d;%dm",
		r, g, b,
	)
}
