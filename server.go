// mad - mock ad server
//   (C) copyright 2016 - J.W. Janssen
package main

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

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
	http.Serve(Listener(), &NoContentHandler{NewLogger()})
}

// EOF
