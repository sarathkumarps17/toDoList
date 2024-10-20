package todo

import (
	"fmt"
	"slices"
)

type task struct {
	id          int
	description string
	complete    bool
	priority    int
}
type tasks []task
type TodoList interface {
	add(task string, priority int)
	list()
	delete(int)
	complete(int)
	getPendingTasks() tasks
	sortByPriority() tasks
}

func (t tasks) add(toDo string, priority int) {
	newTask := task{
		id:          len(t) + 1,
		description: toDo,
		complete:    false,
		priority:    priority,
	}
	t = append(t, newTask)
}

func (t tasks) list() {
	for _, task := range t {
		fmt.Println("Id\t Description\t Complete\t Priority")
		fmt.Printf("%v\t %v\t %v\t %v\n", task.id, task.description, task.complete, task.priority)
	}
}

func (t tasks) delete(id int) {
	for i, task := range t {
		if task.id == id {
			t = append(t[:i], t[i+1:]...)
		}
	}
}

func (t tasks) complete(id int) {
	for i, task := range t {
		if task.id == id {
			task.complete = true
			t[i] = task
		}
	}
}

func (t tasks) getPendingTasks() tasks {
	var pendingTasks tasks

	for _, task := range t {
		if !task.complete {
			pendingTasks = append(pendingTasks, task)
		}
	}

	return pendingTasks
}

func (t tasks) sortByPriority() tasks {
	var sortedTasks tasks
	copy(sortedTasks, t)
	slices.SortFunc(sortedTasks, func(i, j task) int {
		if i.priority > j.priority {
			return 1
		}
		if i.priority < j.priority {
			return -1
		}
		return 0
	})

	return sortedTasks
}
