package cmds

import (
	"fmt"
	"strconv"

	"github.com/joaopedroaat/go-do/internal/services"
	"github.com/spf13/cobra"
)

func DeleteTask(taskService services.TaskService) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				fmt.Println("Invalid id")
				return
			}

			taskService.DeleteTask(id)
		},
	}
}
