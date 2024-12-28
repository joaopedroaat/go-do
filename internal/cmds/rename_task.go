package cmds

import (
	"fmt"
	"strconv"

	"github.com/joaopedroaat/go-do/internal/services"
	"github.com/spf13/cobra"
)

func RenameTask(taskService services.TaskService) *cobra.Command {
	return &cobra.Command{
		Use:   "rename",
		Short: "Rename a task",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				fmt.Println("Invalid id")
				return
			}

			description := args[1]

			taskService.RenameTask(id, description)
		},
	}
}
