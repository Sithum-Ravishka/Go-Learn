package main

import (
	"bufio"
	"fmt"
	"os"
)

// Todo represents a task in the todo list.
type Todo struct {
	Text   string
	Info   string
	IsDone bool
}

// TodoList represents a collection of tasks.
type TodoList struct {
	Tasks []Todo
}

// AddTask adds a new task to the todo list.
func (tl *TodoList) AddTask(task, info string) {
	newTask := Todo{Text: task, Info: info, IsDone: false}
	tl.Tasks = append(tl.Tasks, newTask)
	fmt.Println("Task added:", task)
}

// CompleteTask marks a task as complete.
func (tl *TodoList) CompleteTask(index int) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks[index].IsDone = true
		fmt.Println("Task marked as complete:", tl.Tasks[index].Text)
	} else {
		fmt.Println("Invalid task index.")
	}
}

// UpdateTask updates the text of a task.
func (tl *TodoList) UpdateTask(index int, newText, newInfo string) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks[index].Text = newText
		tl.Tasks[index].Info = newInfo
		fmt.Println("Task updated:", tl.Tasks[index].Text)
	} else {
		fmt.Println("Invalid task index.")
	}
}

// DeleteTask removes a task from the todo list.
func (tl *TodoList) DeleteTask(index int) {
	if index >= 0 && index < len(tl.Tasks) {
		deletedTask := tl.Tasks[index].Text
		tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
		fmt.Println("Task deleted:", deletedTask)
	} else {
		fmt.Println("Invalid task index.")
	}
}

// PrintTodoList displays the current tasks in the todo list.
func (tl *TodoList) PrintTodoList() {
	fmt.Println("Todo List:")
	for i, task := range tl.Tasks {
		status := "Incomplete"
		if task.IsDone {
			status = "Complete"
		}
		fmt.Printf("%d. [%s] %s - %s\n", i+1, status, task.Text, task.Info)
	}
}

func main() {
	todoList := TodoList{}

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Task")
		fmt.Println("2. Complete Task")
		fmt.Println("3. Update Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. View Todo List")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice (1-6): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			task := scanner.Text()

			// Allow users to input additional information
			fmt.Print("Enter additional information: ")
			scanner.Scan()
			info := scanner.Text()

			todoList.AddTask(task, info)
		case 2:
			fmt.Print("Enter task index to mark as complete: ")
			var index int
			_, err := fmt.Scan(&index)
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}
			todoList.CompleteTask(index - 1) // Adjusting for 0-based index
		case 3:
			fmt.Print("Enter task index to update: ")
			var index int
			_, err := fmt.Scan(&index)
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}
			fmt.Print("Enter new task text: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newText := scanner.Text()

			// Allow users to update additional information
			fmt.Print("Enter new additional information: ")
			scanner.Scan()
			newInfo := scanner.Text()

			todoList.UpdateTask(index-1, newText, newInfo) // Adjusting for 0-based index
		case 4:
			fmt.Print("Enter task index to delete: ")
			var index int
			_, err := fmt.Scan(&index)
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}
			todoList.DeleteTask(index - 1) // Adjusting for 0-based index
		case 5:
			todoList.PrintTodoList()
		case 6:
			fmt.Println("Exiting Todo List Program. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
		}
	}
}
