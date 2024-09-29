package task

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	ID    string
	Title string
	Done  bool
}

type TaskList []Task

func NewTask(title string) Task {
	return Task{
		ID:    uuid.New().String(),
		Title: title,
		Done:  false,
	}
}

func (t *Task) Toggle() {
	t.Done = !t.Done
}

func (t TaskList) Print() {
	for _, task := range t {
		if task.Done {
			fmt.Printf("[x] %s\n", task.Title)
		} else {
			fmt.Printf("[ ] %s\n", task.Title)
		}
	}
}

func (t *TaskList) Add(task Task) {
	*t = append(*t, task)
}

func (t *TaskList) Remove(id string) {
	for i, task := range *t {
		if task.ID == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			break
		}
	}
}
