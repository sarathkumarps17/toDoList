package todo

import (
	"fmt"
	"os"
)

func List(t TodoList) {
	t.list()
}

func Add(t TodoList, myTasks tasks) {
	var task string
	var priority int
	fmt.Println("Enter your task")
	fmt.Scanln(&task)
	fmt.Println("Enter your priority")
	fmt.Scanln(&priority)
	t.add(task, priority)
	pendingTasks := myTasks.getPendingTasks()
	fmt.Printf("you have %v tasks pending\n", len(pendingTasks))
}

func Complete(t TodoList) {
	var id int
	fmt.Println("Enter the id of the task you want to complete")
	fmt.Scanln(&id)
	t.complete(id)
	fmt.Println("Successfully completed")
	pendingTasks := t.getPendingTasks()
	fmt.Printf("Now you have %v tasks pending\n", len(pendingTasks))
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

func SortByPriority(t TodoList) tasks {
	return t.sortByPriority()
}

func GetAllCommands(t *TodoList, isVerbose bool) []string {
	c := getInterfaceMethods(t)
	c = append(c, "quit")
	if isVerbose {
		for _, command := range c {
			fmt.Println(command)
		}
		return []string{}
	}

	return c
}

func Start() {
	myTasks := tasks{}
	var myList TodoList
	pendingTasks := myTasks.getPendingTasks()
	fmt.Println("Hi there, Welcome to Go to Do List")
	fmt.Printf("you have %v tasks pending\n", len(pendingTasks))
	fmt.Printf("You can type 'add' to add a new task\n or 'commands' to lst all available commands \n")
	commands := GetAllCommands(&myList, false)

	//read user input
	for {

		var userInput string
		fmt.Scanln(&userInput)
		if !isValidCommand(userInput, commands) {
			fmt.Println("Invalid command")
			continue
		}

		if userInput == "quit" {
			break
		}
		methodInfo := mapMethodToFunction(&myList)

		method, ok := methodInfo[userInput]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		args := []interface{}{}
		args = append(args, myTasks)
		_, err := callMethod(myTasks, method.Name, args...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

}
func isValidCommand(command string, commands []string) bool {
	for _, c := range commands {
		if c == command {
			return true
		}
	}
	return false
}
