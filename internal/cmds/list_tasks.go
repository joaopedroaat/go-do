package cmds

import (
	"fmt"
	"os"

	"github.com/joaopedroaat/go-do/internal/services"
	"github.com/spf13/cobra"
)

func ListTasks(taskService services.TaskService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			completed, err := cmd.Flags().GetBool("completed")
			if err != nil {
				fmt.Println(err)
				return
			}

			listAll, err := cmd.Flags().GetBool("all")
			if err != nil {
				fmt.Println(err)
				return
			}

			if completed && listAll {
				fmt.Println("Error: --completed and --all flags cannot be used together.")
				return
			}

			if listAll {
				taskService.WriteAllTasks(os.Stdout)
			} else if completed {
				taskService.WriteCompletedTasks(os.Stdout)
			} else {
				taskService.WriteNotDoneTasks(os.Stdout)
			}
		},
	}

	cmd.Flags().Bool("completed", false, "List only completed tasks")
	cmd.Flags().Bool("all", false, "List all tasks")

	return cmd
}
