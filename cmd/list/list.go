package list

import (
	"log"

	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "list",
		Short:            "List issues",
		TraverseChildren: true,
		Aliases:          []string{"l"},
		Example:          `  trc l --my`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("list issues")
		},
	}
	return cmd
}
