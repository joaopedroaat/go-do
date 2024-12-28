package main

import (
	"log"

	"github.com/joaopedroaat/go-do/db"
	"github.com/joaopedroaat/go-do/internal/cmds"
	"github.com/joaopedroaat/go-do/internal/services"
	"github.com/spf13/cobra"
)

func main() {
	db, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	taskService := services.NewTaskService(db)

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(
		cmds.ListTasks(taskService),
		cmds.AddTask(taskService),
	)

	rootCmd.Execute()
}
