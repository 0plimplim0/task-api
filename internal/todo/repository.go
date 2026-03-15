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

func (rep *Repository) InsertExamples() {
	rep.lock.Lock()
	defer rep.lock.RUnlock()
	rep.tasks = append(rep.tasks,
		Task{Id: 1, Title: "Task 1", Description: "Example task 1", Completed: false},
		Task{Id: 2, Title: "Task 2", Description: "Example task 2", Completed: false},
		Task{Id: 3, Title: "Task 3", Description: "Example task 3", Completed: false},
	)
}

func (rep *Repository) getTasks() *[]Task {
	rep.lock.RLock()
	defer rep.lock.RUnlock()
	return &rep.tasks
}

func (rep *Repository) getTask(id int) (error, *Task) {
	rep.lock.RLock()
	defer rep.lock.RUnlock()
	for _, v := range rep.tasks {
		if v.Id == id {
			return nil, &v
		}
	}
	return errors.New("There is no Task with the id specified"), nil
}

func (rep *Repository) addTask(task Task) {
	rep.lock.Lock()
	defer rep.lock.Unlock()
	rep.tasks = append(rep.tasks, task)
}

func (rep *Repository) updateTask(id int) error {
	return nil
}

func (rep *Repository) deleteTask(id int) error {
	return nil
}
