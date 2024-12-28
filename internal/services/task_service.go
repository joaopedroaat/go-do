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
	WriteTasks(output io.Writer) error
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
	return nil
}

func (t *taskService) WriteTasks(output io.Writer) error {
	rows, err := t.db.Query("SELECT * FROM Tasks")
	if err != nil {
		return err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Description, &task.Done)
		if err != nil {
			return err
		}

		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return nil
	}

	fmt.Println(len(tasks))

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Id\tDescription\tStatus")
	for _, t := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%t\n", t.Id, t.Description, t.Done)
	}

	w.Flush()

	return nil
}
