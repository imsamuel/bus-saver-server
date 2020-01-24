package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleIncomingBusesGet(t *testing.T) {
	// Create a Request to pass to handler we want to test.
	req, err := http.NewRequest("GET", "/incoming-buses/97331/4", nil)
	if err != nil {
		panic(err)
	}

	// Create a ResponseRecorder (satisfies the http.ResponseWriter interface)
	// to record the response.
	rr := httptest.NewRecorder()

	// Create a server and initialize its depdencies (handleIncomingBusesGet)
	// is a method off of it.
	srv := server{
		router:httprouter.New(),
	}
	srv.routes()

	// Will call handleIncomingBuses.
	srv.router.ServeHTTP(rr, req)

	// Tests for a 200 status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code of %d but got %d", http.StatusOK, status)
	}
}
