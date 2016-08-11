// mad - mock ad server
//   (C) copyright 2016 - J.W. Janssen
package main

import (
	"flag"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
)

var debug bool
var bindAddr string

func init() {
	flag.BoolVar(&debug, "debug", false, "Enable debugging mode (do not use Socket activation on Linux).")
	flag.StringVar(&bindAddr, "bind", ":8080", "Set address and port to bind to.")
}

type NoContentHandler struct {
	log Logger
}

type Logger interface {
	Log(msg string, args ...interface{})
}

type StdErrLogger struct {
}

func (l *StdErrLogger) Log(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (h *NoContentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Log("Request from %s on %s for %s.", r.RemoteAddr, r.URL, r.Host)

	ctype := mime.TypeByExtension(filepath.Ext(r.URL.Path))
	if ctype == "" {
		ctype = "text/plain"
	}

	w.Header().Add("Content-Type", ctype)
	w.Header().Add("Connection", "close")
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	flag.Parse()

	http.Serve(Listener(), &NoContentHandler{NewLogger()})
}

func defaultListener() net.Listener {
	ln, err := net.Listen("tcp", bindAddr)
	if err != nil {
		panic(err)
	}
	return ln
}

// EOF
