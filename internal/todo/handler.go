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

func GetAllTasks(w http.ResponseWriter, r *http.Request, rep *Repository) {
	tasks := getTasks(rep)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request, rep *Repository) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err, task := getTask(id, rep)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func AddTask(w http.ResponseWriter, r *http.Request, rep *Repository) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	addTask(task, rep)
	w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request, rep *Repository) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request, rep *Repository) {

}
