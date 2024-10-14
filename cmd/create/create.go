package create

import (
	"log"

	"github.com/fgazat/trc/client"
	"github.com/fgazat/trc/config"
	"github.com/spf13/cobra"
)

func Create(cfg *config.Config, creator client.Creator) *cobra.Command {
	var queue, summary, description string
	cmd := &cobra.Command{
		Use:              "create",
		Short:            "Create issue",
		TraverseChildren: true,
		Aliases:          []string{"c"},
		Example:          `  trc c -q "TEST" -m "hello" -d "world"`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Creating issue")
			key, err := creator.CreateIssue(queue, summary, description)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf(cfg.WebBaseURL + "/" + key)
		},
	}
	cmd.Flags().StringVarP(&queue, "queue", "q", cfg.Issues.DefaultQueue, "Queue")
	cmd.Flags().StringVarP(&summary, "sum", "m", "", "Summary. Short -m is due to convinient commit message arg")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "Description of issue")
	return cmd
}
