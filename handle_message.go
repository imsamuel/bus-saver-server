package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *server) handleMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse the request body then extract the values.
	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	message := r.Form.Get("message")

	// Get a Message instance and populate with posted values.
	m := Message{name, email, message}

	// Return a response of status code 502 if failure to connect to the
	// Gmail SMTP server, else return a response of status code 200.
	if err := sendEmail(m); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
