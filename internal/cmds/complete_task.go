package cmds

import (
	"fmt"
	"strconv"

	"github.com/joaopedroaat/go-do/internal/services"
	"github.com/spf13/cobra"
)

func CompleteTask(taskService services.TaskService) *cobra.Command {
	return &cobra.Command{
		Use:   "complete [id]",
		Short: "Mark a task as completed",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				fmt.Println("Invalid id")
				return
			}

			taskService.CompleteTask(id)
		},
	}
}
