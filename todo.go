package todo

import (
	"bufio"
	"fmt"
	"os"
)

func List(t TodoList) {
	t.list()
}

func Add(t TodoList) {
	var task string
	var priority int

	// Create a scanner to read the full line for task
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your task")
	if scanner.Scan() {
		task = scanner.Text()
	}

	fmt.Println("Enter your priority")
	_, err := fmt.Scanf("%d\n", &priority)
	if err != nil {
		fmt.Println("Please enter a valid number for priority")
		return
	}

	t.add(task, priority)
	pendingTasks := t.getPendingTasks()
	fmt.Printf("you have %v tasks pending\n", len(pendingTasks))
	List(t)
}

func Complete(t TodoList) {
	var id int
	fmt.Println("Enter the id of the task you want to complete")
	fmt.Scanln(&id)
	t.complete(id)
	fmt.Println("Successfully completed")
	pendingTasks := t.getPendingTasks()
	fmt.Printf("Now you have %v tasks pending\n", len(pendingTasks))
	List(t)
}

func Delete(t TodoList) {
	var id int
	fmt.Println("Enter the id of the task you want to delete")
	fmt.Scanln(&id)
	t.delete(id)
	fmt.Println("Successfully deleted")
	pendingTasks := t.getPendingTasks()
	fmt.Printf("Now you have %v tasks pending\n", len(pendingTasks))

}

func SortByPriority(t TodoList) {
	t.sortByPriority()
}

func createCommandMap(t TodoList) map[string]func() {
	commands := map[string]func(){
		"list": func() {
			List(t)
		},
		"add": func() {
			Add(t)
		},
		"complete": func() {
			Complete(t)
		},
		"delete": func() {
			Delete(t)
		},
		"sort": func() {
			SortByPriority(t)
		},
		"quit": func() {
			fmt.Println("Bye..")
			os.Exit(0)
		},
	}

	return commands
}

func Start() {
	var myTasks tasks
	var todoList TodoList = &myTasks
	pendingTasks := myTasks.getPendingTasks()
	fmt.Println("Hi there, Welcome to Go to Do List")
	fmt.Printf("you have %v tasks pending\n", len(pendingTasks))

	//read user input
	for {
		fmt.Printf("You can type 'add' to add a new task or 'help' to lst all available commands \n")
		var userInput string
		fmt.Scanln(&userInput)
		commandMap := createCommandMap(todoList)
		if userInput == "help" {
			for key := range commandMap {
				fmt.Println(key)
			}
			continue
		}
		fn, exists := commandMap[userInput]
		if !exists {
			fmt.Println("Sorry invalid command")
			continue
		}
		fn()
	}

}
