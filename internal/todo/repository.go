package todo

import (
	"errors"
	"sync"
)

type Repository struct {
	tasks []Task
	lock  sync.RWMutex
}

func InitRepository() *Repository {
	return &Repository{}
}

func InsertExamples(rep *Repository) {
	rep.lock.Lock()
	defer rep.lock.RUnlock()
	rep.tasks = append(rep.tasks,
		Task{Id: 1, Title: "Task 1", Description: "Example task 1", Completed: false},
		Task{Id: 2, Title: "Task 2", Description: "Example task 2", Completed: false},
		Task{Id: 3, Title: "Task 3", Description: "Example task 3", Completed: false},
	)
}

func getTasks(rep *Repository) *[]Task {
	rep.lock.RLock()
	defer rep.lock.RUnlock()
	return &rep.tasks
}

func getTask(id int, rep *Repository) (error, *Task) {
	rep.lock.RLock()
	defer rep.lock.RUnlock()
	for _, v := range rep.tasks {
		if v.Id == id {
			return nil, &v
		}
	}
	return errors.New("There is no Task with the id specified"), nil
}

func addTask(task Task, rep *Repository) {
	rep.lock.Lock()
	defer rep.lock.Unlock()
	rep.tasks = append(rep.tasks, task)
}

func updateTask(id int, rep *Repository) error {
	return nil
}

func deleteTask(id int, rep *Repository) error {
	return nil
}
