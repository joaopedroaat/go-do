package main

import (
	"os"

	"github.com/joaopedroaat/go-do/internal/models"
)

func main() {
	models.AddTask("Take dog for a walk")
	models.AddTask("Wash the dishes")

	models.CompleteTask(2)

	models.AddTask("Take drugs")

	models.DeleteTask(3)

	models.AddTask("Buy toilet paperrr")

	models.RenameTask(4, "Buy toilet paper")

	models.WriteTasks(os.Stdout)
}
