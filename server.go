// mad - mock ad server
//   (C) copyright 2015 - J.W. Janssen
package main

import (
	"net/http"
)

type NoContentHandler struct {
	log Logger
}

type Logger interface {
	Log(msg string, args ...interface{})
}

func (h *NoContentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Log("Request from %s on %s for %s.", r.RemoteAddr, r.URL, r.Host)

	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Connection", "close")
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.Serve(Listener(), &NoContentHandler{NewLogger()})
}

// EOF
