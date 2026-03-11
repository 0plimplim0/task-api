package main

import (
	"net/http"
	"task-api/internal/todo"
)

func setRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", todo.Greeting)
	mux.HandleFunc("GET /api/test", todo.JsonTest)

	return mux
}
