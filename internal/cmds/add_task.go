package cmds

import (
	"github.com/joaopedroaat/go-do/internal/services"
	"github.com/spf13/cobra"
)

func AddTask(taskService services.TaskService) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Create new todo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskService.AddTask(args[0])
		},
	}
}
