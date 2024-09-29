package storage

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"todo-go/task"
)

func SaveTasks(tasks task.TaskList) error {
	csvFile, err := os.Create("tasks.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, task := range tasks {
		writer.Write([]string{task.ID, task.Title, task.Description, strconv.FormatBool(task.Done)})
	}

	return nil
}

func LoadTasks() (task.TaskList, error) {
	file, err := os.Open("tasks.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return task.TaskList{}, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var tasks task.TaskList
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task.Task{
			ID:          record[0],
			Title:       record[1],
			Description: record[2],
			Done:        record[3] == "true",
		})
	}

	return tasks, nil
}
