package main

import (
	"net/http"
	"task-api/internal/todo"
)

func setRouter(rep *todo.Repository) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", todo.Greeting)
	mux.HandleFunc("GET /tasks", todo.GetAllTasks)
	mux.HandleFunc("GET /tasks/{id}", todo.GetTaskById)
	mux.HandleFunc("POST /tasks", todo.AddTask)
	mux.HandleFunc("PUT /tasks/{id}", todo.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", todo.DeleteTask)

	return mux
}
