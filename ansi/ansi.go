package ansi

import "fmt"

type Code int

func (c Code) String() string {
	return fmt.Sprintf("\x1b[%dm", c)
}

// this function resets all the previous ansi codes,
// cast is needed to use with any Code aliases
func Apply(str string, codes ...Code) string {
	var prefix string

	for _, code := range codes {
		if code != 0 {
			prefix += code.String()
		}
	}

	return fmt.Sprintf("%s%s%s", prefix, str, ResetAll)
}
