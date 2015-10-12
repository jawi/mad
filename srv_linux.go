// mad - mock ad server
//   (C) copyright 2015 - J.W. Janssen
package main

import (
	"net"
	"fmt"

	"github.com/coreos/go-systemd/activation"
	"github.com/coreos/go-systemd/journal"
)

type JournaldLogger struct {
}

func (l *JournaldLogger) Log(msg string, args ...interface{}) {
	if journal.Enabled() {
	    journal.Print(journal.PriInfo, fmt.Sprintf(msg, args...))
	}
}

func NewLogger() Logger {
	return &JournaldLogger{}
}

func Listener() net.Listener {
	listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}

	if len(listeners) != 1 {
		panic("Unexpected number of socket activation fds")
	}

	return listeners[0]
}

// EOF
