package models

import (
	"fmt"
	"io"
	"text/tabwriter"
)

type task struct {
	id          uint64
	description string
	done        bool
}

var (
	tasks     []task
	idCounter uint64
)

func AddTask(description string) {
	idCounter++
	t := task{
		id:          idCounter,
		description: description,
		done:        false,
	}

	tasks = append(tasks, t)
}

func CompleteTask(id uint64) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].id == id {
			tasks[i].done = true
		}
	}
}

func DeleteTask(id uint64) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}

func WriteTasks(output io.Writer) {
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Id\tDescription\tStatus")
	for _, t := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%t\n", t.id, t.description, t.done)
	}

	w.Flush()
}
