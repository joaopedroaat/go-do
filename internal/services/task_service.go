package services

import (
	"database/sql"
	"io"
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
	return nil
}

func (t *taskService) CompleteTask(id uint64) error {
	return nil
}

func (t *taskService) RenameTask(id uint64, description string) error {
	return nil
}

func (t *taskService) DeleteTask(id uint64) error {
	return nil
}

func (t *taskService) WriteTasks(output io.Writer) error {
	return nil
}
