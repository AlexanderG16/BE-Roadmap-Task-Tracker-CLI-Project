package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// TODO: Make a CLI App for Task Tracking
// Tasks are saved in a JSON file
// Make a command to add a task
// a command to list tasks

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Status      string `json:"status"` // "todo", "in progress", "done"
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func loadTasks() ([]Task, error) {
	if _, err := os.Stat("./tasks/tasks.json"); os.IsNotExist(err) {
		return nil, fmt.Errorf("tasks.json file does not exist, please create the file first")
	}
	file, err := os.Open("./tasks/tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTask(task Task) {
	var tasks []Task
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	tasks = append(tasks, task)

	file, err := os.OpenFile("./tasks/tasks.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [add|update|delete|mark|list]")
		return
	}

	switch os.Args[1] {
	case "add":
		result := strings.Split(os.Args[2], "<>")
		createdAt := time.Now().Format(time.RFC3339)
		id := strings.Join([]string{time.Now().Format("0102T15:04"), strings.ToUpper(result[0][0:2])}, "-")
		status := "todo"
		task := Task{
			ID:          id,
			Description: result[0],
			Group:       result[1],
			Status:      status,
			CreatedAt:   createdAt,
			UpdatedAt:   createdAt,
		}
		saveTask(task)
		fmt.Println("Task added successfully (ID: ", task.ID, ")")
	case "update":
		// TODO: Implement update task
	case "delete":
		// TODO: Implement delete task
	case "mark":
		// TODO: Implement mark task
	case "list":
		// TODO: Implement list tasks
		tasks, err := loadTasks()
		if err != nil {
			panic(err)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			fmt.Println("Use 'add' command to create a new task.")
			return
		}

		if len(os.Args) == 2 {
			// TODO: Implement list all tasks
			for _, task := range tasks {
				fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
					task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
				fmt.Printf("------------------------------------------\n")
			}
			return
		}

		switch os.Args[2] {
		case "todo":
			// TODO: Implement list todo tasks
		case "in-progress":
			// TODO: Implement list in progress tasks
		case "done":
			// TODO: Implement list done tasks
		case "group":
			// TODO: Implement list grouped tasks
		default:
			fmt.Println("Unknown list option: ", os.Args[2])
		}
	default:
		fmt.Println("Unknown command: ", os.Args[1])
	}
}
