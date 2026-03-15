package main

import (
	"net/http"
	"task-api/internal/todo"
)

func setRouter(rep *todo.Repository) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", todo.Greeting)
	mux.HandleFunc("GET /tasks", todo.GetAllTasks(rep))
	mux.HandleFunc("GET /tasks/{id}", todo.GetTaskById(rep))
	mux.HandleFunc("POST /tasks", todo.AddTask(rep))
	mux.HandleFunc("PUT /tasks/{id}", todo.UpdateTask(rep))
	mux.HandleFunc("DELETE /tasks/{id}", todo.DeleteTask(rep))

	return mux
}
