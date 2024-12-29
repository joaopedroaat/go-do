package services

import (
	"database/sql"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/joaopedroaat/go-do/internal/models"
)

type taskService struct {
	db *sql.DB
}

type TaskService interface {
	AddTask(description string) error
	CompleteTask(id uint64) error
	RenameTask(id uint64, description string) error
	DeleteTask(id uint64) error
	WriteAllTasks(output io.Writer) error
	WriteCompletedTasks(output io.Writer) error
	WriteNotDoneTasks(output io.Writer) error
}

func NewTaskService(db *sql.DB) *taskService {
	return &taskService{
		db: db,
	}
}

func (t *taskService) AddTask(description string) error {
	query := "INSERT INTO Tasks (description) VALUES (?)"
	_, err := t.db.Exec(query, description)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskService) CompleteTask(id uint64) error {
	query := "UPDATE Tasks SET done = ? WHERE id = ?"
	_, err := t.db.Exec(query, true, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskService) RenameTask(id uint64, description string) error {
	query := "UPDATE Tasks SET description = ? WHERE id = ?"
	_, err := t.db.Exec(query, description, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskService) DeleteTask(id uint64) error {
	query := "DELETE FROM Tasks WHERE id = ?"
	_, err := t.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskService) WriteAllTasks(output io.Writer) error {
	rows, err := t.db.Query("SELECT * FROM Tasks")
	if err != nil {
		return err
	}

	return writeTasks(output, rows)
}

func (t *taskService) WriteCompletedTasks(output io.Writer) error {
	rows, err := t.db.Query("SELECT * FROM Tasks WHERE done = TRUE")
	if err != nil {
		return err
	}

	return writeTasks(output, rows)
}

func (t *taskService) WriteNotDoneTasks(output io.Writer) error {
	rows, err := t.db.Query("SELECT * FROM Tasks WHERE done = FALSE")
	if err != nil {
		return err
	}

	return writeTasks(output, rows)
}

func writeTasks(output io.Writer, tasks *sql.Rows) error {
	defer tasks.Close()

	var tasksSlice []models.Task

	for tasks.Next() {
		var task models.Task
		err := tasks.Scan(&task.Id, &task.Description, &task.Done)
		if err != nil {
			return err
		}

		tasksSlice = append(tasksSlice, task)
	}

	if len(tasksSlice) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Id\tDescription\tStatus")
	for _, t := range tasksSlice {
		fmt.Fprintf(w, "%d\t%s\t%t\n", t.Id, t.Description, t.Done)
	}

	w.Flush()

	return nil
}
