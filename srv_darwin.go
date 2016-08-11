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
	return defaultListener()
}

// EOF
