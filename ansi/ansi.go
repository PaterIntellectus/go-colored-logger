package ansi

import (
	"fmt"
)

type Code int

func (c Code) String() string {
	return fmt.Sprintf("\x1b[%dm", c)
}

// this function resets all the previous ansi codes,
// cast is needed to use with any Code aliases
func Apply(str string, codes ...Code) string {
	length := len(codes)

	if length < 1 {
		return str
	}

	prefix := fmt.Sprintf("\033[%d", codes[0])

	for i := 1; i < length; i++ {
		code := codes[i]
		if code != 0 {
			prefix = fmt.Sprintf("%s;%d", prefix, code)
		}
	}

	prefix = fmt.Sprintf("%sm", prefix)

	return fmt.Sprintf("%s%s%s", prefix, str, ResetAll)
}
