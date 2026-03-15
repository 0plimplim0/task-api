package main

import (
	"fmt"
	"net/http"
	"task-api/internal/todo"
	"time"
)

func main() {

	rep := todo.InitRepository()

	mux := setRouter(rep)
	rep.InsertExamples()

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
