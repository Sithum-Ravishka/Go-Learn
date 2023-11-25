package main

import (
	"fmt"
	"os"
)

// Todo represents a task in the todo list.
type Todo struct {
	Text   string
	IsDone bool
}

// TodoList represents a collection of tasks.
type TodoList struct {
	Tasks []Todo
}

// AddTask adds a new task to the todo list.
func (tl *TodoList) AddTask(task string) {
	newTask := Todo{Text: task, IsDone: false}
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
func (tl *TodoList) UpdateTask(index int, newText string) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks[index].Text = newText
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
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
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
			var task string
			fmt.Scan(&task)
			todoList.AddTask(task)
		case 2:
			fmt.Print("Enter task index to mark as complete: ")
			var index int
			fmt.Scan(&index)
			todoList.CompleteTask(index - 1)
		case 3:
			fmt.Print("Enter task index to update: ")
			var index int
			fmt.Scan(&index)
			fmt.Print("Enter new task text: ")
			var newText string
			fmt.Scan(&newText)
			todoList.UpdateTask(index-1, newText)
		case 4:
			fmt.Print("Enter task index to delete: ")
			var index int
			fmt.Scan(&index)
			todoList.DeleteTask(index - 1)
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
