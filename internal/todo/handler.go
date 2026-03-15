package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi! Welcome to my server!")
}

func GetAllTasks(rep *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks := rep.getTasks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTaskById(rep *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		err, task := rep.getTask(id)
		if err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

func AddTask(rep *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		rep.addTask(task)
		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateTask(rep *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		var task Task
		jsonErr := json.NewDecoder(r.Body).Decode(&task)
		if jsonErr != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		err = rep.updateTask(id, task)
		if err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTask(rep *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		err = rep.deleteTask(id)
		if err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// TODOS:
// - Add validation for empty json in post
// - Add validation for empty fields
// - Add validation for POST to existing ids
// - Add validation for Content-Type (if not then 415 Unsupported Media Type)
// - Add DoS protection with http.MaxBytesReader
