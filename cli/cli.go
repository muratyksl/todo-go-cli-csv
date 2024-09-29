package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-go/storage"
	"todo-go/task"
)

func Run() {
	taskList, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	fmt.Println("Welcome to todo-go")
	fmt.Println("Type 'help' to see available commands")
	fmt.Println("Type 'exit' to exit")

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch {
		case text == "exit":
			err := storage.SaveTasks(taskList)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
			os.Exit(0)
		case text == "help":
			printHelp()
		case strings.HasPrefix(text, "add "):
			parts := strings.SplitN(strings.TrimPrefix(text, "add "), "|", 2)
			title := strings.TrimSpace(parts[0])
			description := ""
			if len(parts) > 1 {
				description = strings.TrimSpace(parts[1])
			}
			newTask := task.NewTask(title, description)
			taskList.Add(newTask)
			fmt.Printf("Task added: %s | %s\n", newTask.Title, newTask.Description)
			err := storage.SaveTasks(taskList)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		case strings.HasPrefix(text, "remove "):
			id := strings.TrimPrefix(text, "remove ")
			taskList.Remove(id)
			err := storage.SaveTasks(taskList)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		case strings.HasPrefix(text, "toggle "):
			id := strings.TrimPrefix(text, "toggle ")
			toggleTask(&taskList, id)
			err := storage.SaveTasks(taskList)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		case strings.HasPrefix(text, "list"):
			taskList.Print()
		case strings.HasPrefix(text, "clear"):
			taskList = task.TaskList{}
			err := storage.SaveTasks(taskList)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add <title>|<description>  - add a new task")
	fmt.Println("  remove <id>                - remove a task")
	fmt.Println("  toggle <id>                - toggle a task")
	fmt.Println("  list                       - list all tasks")
	fmt.Println("  clear                      - clear all tasks")
	fmt.Println("  exit                       - exit the application")
}

func toggleTask(taskList *task.TaskList, id string) {
	for i := range *taskList {
		if (*taskList)[i].ID == id {
			(*taskList)[i].Toggle()
			break
		}
	}
}
