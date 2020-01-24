package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	router *httprouter.Router
}

// ServeHTTP calls the router's ServeHTTP method - allows server to implement
// the http.Handler interface.
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
