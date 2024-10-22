package todo

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"
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
	sortByPriority()
}

func (t *tasks) add(toDo string, priority int) {
	newTask := task{
		id:          len(*t) + 1,
		description: toDo,
		complete:    false,
		priority:    priority,
	}
	*t = append(*t, newTask)
}

func (t tasks) list() {
	// Initialize tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Print header
	fmt.Fprintln(w, "ID\tDescription\tComplete\tPriority\t")
	fmt.Fprintln(w, "----\t-----------\t--------\t--------\t")

	// Print tasks
	for _, task := range t {
		fmt.Fprintf(w, "%d\t%s\t%v\t%d\t\n",
			task.id,
			task.description,
			task.complete,
			task.priority)
	}

	// Flush the writer to display the output
	w.Flush()
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

func (t tasks) sortByPriority() {
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

	sortedTasks.list()
}
