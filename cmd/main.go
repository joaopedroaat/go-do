package main

import (
	"fmt"
	"os"

	"github.com/joaopedroaat/go-do/internal/models"
)

func main() {
	fmt.Println("Hello, world!")

	models.AddTask("Take dog for a walk")
	models.AddTask("Wash the dishes")

	models.WriteTasks(os.Stdout)
}
