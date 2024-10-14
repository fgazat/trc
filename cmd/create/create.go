package create

import (
	"log"

	"github.com/spf13/cobra"
)

func Create() *cobra.Command {
	var summary string
	var description string
	cmd := &cobra.Command{
		Use:              "create",
		Short:            "Create issue",
		TraverseChildren: true,
		Aliases:          []string{"c"},
		Example:          `  trc c -m "hello" -d "world"`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("create issue")
		},
	}
	cmd.Flags().StringVarP(&summary, "sum", "m", "", "Summary. Short -m is due to convinient commit message arg")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "Description of issue")
	return cmd
}
