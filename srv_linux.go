// mad - mock ad server
//   (C) copyright 2016 - J.W. Janssen
package main

import (
	"fmt"
	"net"

	"github.com/coreos/go-systemd/activation"
	"github.com/coreos/go-systemd/journal"
)

type JournaldLogger struct {
}

func (l *JournaldLogger) Log(msg string, args ...interface{}) {
	journal.Print(journal.PriInfo, msg, args...)
}

func NewLogger() Logger {
	if journal.Enabled() {
		return &JournaldLogger{}
	}
	return &StdErrLogger{}
}

func Listener() net.Listener {
	listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}

	if len(listeners) != 1 {
		panic(fmt.Sprintf("Unexpected number of socket activation fds, got: %d listeners, expected 1!", len(listeners)))
	}

	return listeners[0]
}

// EOF
