package main

import (
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHandleMessage tests the handleMessage handler.
// Expects handleMessage to:
// - respond with a status code of 200 in the case of email being successfully
//   sent. Else, respond with a status code of 502 (Bad Gateway Error) because
//   this server has run into trouble trying to connect to Gmail's STMP server.
func TestHandleMessage(t *testing.T) {
	// Dummy message to use as a payload for sendMail.
	m := Message{
		Name:    "Snoop Dogg",
		Email:   "snoopdogg@weedmail.com",
		Message: "Smoke weed everyday",
	}
	// Convert the message to an io.Reader to satisfy NewRequest's body param.
	mBytes, err := json.Marshal(m) // to slice of bytes
	if err != nil {
		panic(err)
	}
	mReader := bytes.NewReader(mBytes) // to io.Reader

	// Create a request with the dummy message to pass to handleMessage.
	req, err := http.NewRequest("POST", "/message", mReader)

	// Create a ResponseRecorder to record the response - satifies the
	// http.ResponseWriter interface so can be passed to as the second argument
	// to a http.HandlerFunc.
	rr := httptest.NewRecorder()

	// Set up server and its dependencies - need to do this because
	// handleMessage is method hanging off of it.
	srv := server{router:httprouter.New()}
	srv.routes() // configure routes

	// Call handleMessage with the Request and ResponseRecorder.
	srv.handleMessage()(rr, req)

	// Check that the status code is 200.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %d want %d", status, http.StatusOK)
	}
}
