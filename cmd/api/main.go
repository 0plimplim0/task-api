package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	mux := setRouter()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Server running on http://localhost%s\n", srv.Addr)
	srv.ListenAndServe()
}
