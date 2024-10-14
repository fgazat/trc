package update

import (
	"log"

	"github.com/spf13/cobra"
)

func Update() *cobra.Command {
	var summary string
	var description string
	cmd := &cobra.Command{
		Use:              "update",
		Short:            "Update issue",
		TraverseChildren: true,
		Aliases:          []string{"u"},
		Example:          `  trc u -m "hello" -d "world"`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("update issue")
		},
	}
	cmd.Flags().StringVarP(&summary, "sum", "m", "", "Summary. Short -m is due to convinient commit message arg")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "Description of issue")
	return cmd
}
