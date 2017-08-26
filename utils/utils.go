package utils

import (
	"fmt"
	"os"
)

func Check(err error, message string) {
	if err == nil {
		return
	}

	if message == "" {
		message = err.Error()
	}

	Exit(message)
}

func Exit(a ...interface{}) {
	Errorln(a...)
	os.Exit(1)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

func Errorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}
