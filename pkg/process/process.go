package process

import (
	"os"
)

var osExit = os.Exit

// Bail terminates program execution with an error message
func Bail(message string, err error) {
	code := 0
	if err != nil {
		os.Stderr.WriteString(message + `: ` + err.Error() + "\n")
		code = 1
	}
	osExit(code)
}
