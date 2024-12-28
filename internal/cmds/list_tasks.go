package cmds

import (
	"os"

	"github.com/joaopedroaat/go-do/internal/services"
	"github.com/spf13/cobra"
)

func ListTasks(taskService services.TaskService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			taskService.WriteTasks(os.Stdout)
		},
	}
}
