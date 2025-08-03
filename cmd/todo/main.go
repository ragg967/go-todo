package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	name      string
	task      string
	startDate string
	endDate   string
	isDone    bool // default is false for bool type
}

func NewTask() *Task {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the name of the task: ")
	scanner.Scan()
	n := strings.TrimSpace(scanner.Text())
	if n == "" {
		fmt.Println("Task name cannot be empty")
		return nil
	}

	fmt.Print("Enter the task description: ")
	scanner.Scan()
	t := strings.TrimSpace(scanner.Text())
	if t == "" {
		fmt.Println("Task description cannot be empty")
		return nil
	}

	var sd string = time.Now().Format("2006-01-02")

	fmt.Print("Enter the end date (YYYY-MM-DD): ")
	scanner.Scan()
	ed := strings.TrimSpace(scanner.Text())
	if ed == "" {
		fmt.Println("End date cannot be empty")
		return nil
	}

	return &Task{
		name:      n,
		task:      t,
		startDate: sd,
		endDate:   ed,
		isDone:    false,
	}
}

func EditTask(task *Task) *Task {
	if task == nil {
		fmt.Println("No task provided to edit")
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)

	var options string = `What do you want to edit?
N: name
T: task
E: end date`
	println(options)

	fmt.Print("Choose an option: ")
	scanner.Scan()
	choiceStr := strings.TrimSpace(strings.ToLower(scanner.Text()))

	if len(choiceStr) == 0 {
		fmt.Println("No option selected")
		return task
	}

	chosenOption := rune(choiceStr[0])

	switch chosenOption {
	case 'n':
		fmt.Print("What should the new name be: ")
		scanner.Scan()
		newName := strings.TrimSpace(scanner.Text())
		if newName != "" {
			task.name = newName
		} else {
			fmt.Println("Name cannot be empty")
		}

	case 't':
		fmt.Print("What should the new task description be: ")
		scanner.Scan()
		newTask := strings.TrimSpace(scanner.Text())
		if newTask != "" {
			task.task = newTask
		} else {
			fmt.Println("Task description cannot be empty")
		}

	case 'e':
		fmt.Print("What should the new end date be (YYYY-MM-DD): ")
		scanner.Scan()
		newEndDate := strings.TrimSpace(scanner.Text())
		if newEndDate != "" {
			task.endDate = newEndDate
		} else {
			fmt.Println("End date cannot be empty")
		}

	default:
		fmt.Println("Invalid option selected")
		return task
	}

	fmt.Println("Task updated successfully!")
	return task
}

// Helper function to display a task
func (t *Task) Display() {
	status := "Not Done"
	if t.isDone {
		status = "Done"
	}
	fmt.Printf("Task: %s\nDescription: %s\nStart Date: %s\nEnd Date: %s\nStatus: %s\n",
		t.name, t.task, t.startDate, t.endDate, status)
}

// Example usage
func main() {
	// Create a new task
	fmt.Println("=== Creating a new task ===")
	task := NewTask()
	if task != nil {
		fmt.Println("\nTask created:")
		task.Display()

		// Edit the task
		fmt.Println("\n=== Editing the task ===")
		EditTask(task)
		fmt.Println("\nUpdated task:")
		task.Display()
	}
}
