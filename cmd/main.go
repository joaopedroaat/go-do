package main

import (
	"log"
	"os"

	"github.com/joaopedroaat/go-do/db"
	"github.com/joaopedroaat/go-do/internal/services"
)

func main() {
	db, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	taskService := services.NewTaskService(db)

	taskService.AddTask("Take dog for a walk")
	taskService.AddTask("Wash the dishes")

	taskService.CompleteTask(40)

	taskService.RenameTask(40, "Go to gym")

	taskService.DeleteTask(30)
	taskService.DeleteTask(31)
	taskService.DeleteTask(32)
	taskService.DeleteTask(33)
	taskService.DeleteTask(34)
	taskService.DeleteTask(35)
	taskService.DeleteTask(36)
	taskService.DeleteTask(37)
	taskService.DeleteTask(38)
	taskService.DeleteTask(39)

	taskService.WriteTasks(os.Stdout)
}
