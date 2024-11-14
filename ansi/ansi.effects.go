package ansi

type Effect Code

const (
	Bold         Effect = 1
	Faint        Effect = 2
	Italic       Effect = 3
	Underline    Effect = 4
	BlinkSlow    Effect = 5
	BlinkFast    Effect = 6
	ReverseVideo Effect = 7
	Conceal      Effect = 8
	CrossedOut   Effect = 9

	Fraktur         Effect = 20
	DoubleUnderline Effect = 21

	Framed    Effect = 51
	Encircled Effect = 52
	Overlined Effect = 53
)
