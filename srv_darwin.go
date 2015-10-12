// mad - mock ad server
//   (C) copyright 2015 - J.W. Janssen
package main

import (
	"fmt"
	"net"
	"os"
)

type StdErrLogger struct {
}

func (l *StdErrLogger) Log(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf(msg, args))
	fmt.Fprintf(os.Stderr, "\n")
}

func NewLogger() Logger {
	return &StdErrLogger{}
}

func Listener() net.Listener {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	return ln
}

// EOF
