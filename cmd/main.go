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

	taskService.WriteTasks(os.Stdout)
}
