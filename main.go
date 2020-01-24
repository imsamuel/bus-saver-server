package main

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

// Handles any error returned from function run.
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// Called from function main.
func run() error {
	if os.Getenv("ACCOUNT_KEY") == "" {
		return errors.New("account key for API has to be set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	srv := server{
		router:httprouter.New(),
	}
	srv.routes()
	if err := http.ListenAndServe(":" + port, srv); err != nil {
		return err
	}

	return nil
}
