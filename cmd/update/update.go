package update

import (
	"log"

	"github.com/fgazat/trc/config"
	"github.com/spf13/cobra"
)

func Update(cfg *config.Config) *cobra.Command {
	var summary, key, description string
	cmd := &cobra.Command{
		Use:              "update",
		Short:            "Update issue",
		TraverseChildren: true,
		Aliases:          []string{"u"},
		Example:          `  trc u -k "TEST-1" -m "hello" -d "world"`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("update issue")
		},
	}
	cmd.Flags().StringVarP(&key, "key", "k", "", "Issue key to update")
	cmd.Flags().StringVarP(&summary, "sum", "m", "", "Summary. Short -m is due to convinient commit message arg")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "Description of issue")
	return cmd
}
