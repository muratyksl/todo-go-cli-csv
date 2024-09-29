package task

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Done        bool
}

type TaskList []Task

func NewTask(title string, description string) Task {
	return Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Done:        false,
	}
}

func (t *Task) Toggle() {
	t.Done = !t.Done
}

func (t TaskList) Print() {
	for _, task := range t {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%s %s: %s\n", status, task.Title, task.Description)
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
