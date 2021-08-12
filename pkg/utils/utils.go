package utils

import (
	"fmt"
	"io"
	"os"
)

// Info prints to screen
func Info(stdout io.Writer, format string, args ...interface{}) {
	fmt.Fprintf(stdout, "\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// CheckIfError handles errors and exits if necessary
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
