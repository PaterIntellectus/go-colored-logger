package main

import (
	"github.com/PaterIntellectus/go-colored-logger/logger"
)

func main() {
	logger := logger.New(
		&logger.Options{Stdout: true, Filename: ""},
	)

	prefix := "This is:"

	logger.I("%s Info", prefix)
	logger.W("%s Warning", prefix)
	logger.E("%s Error", prefix)
	logger.F("%s Fatal", prefix)

	logger.I("%s after Fatal", prefix)
}
