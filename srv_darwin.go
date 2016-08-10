// mad - mock ad server
//   (C) copyright 2016 - J.W. Janssen
package main

import (
	"net"
)

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
