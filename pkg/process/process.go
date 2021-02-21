package process

import (
	"os"
)

var osExit = os.Exit

// Bail terminates program execution with an error message
func Bail(message string, err error) {
	if err != nil {
		os.Stderr.WriteString(message + `: ` + err.Error() + "\n")
		osExit(1)
	} else {
		osExit(0)
	}
}
