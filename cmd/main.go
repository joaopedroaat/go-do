package main

import (
	"os"

	"github.com/joaopedroaat/go-do/internal/models"
)

func main() {
	models.AddTask("Take dog for a walk")
	models.AddTask("Wash the dishes")

	models.CompleteTask(2)

	models.WriteTasks(os.Stdout)
}
