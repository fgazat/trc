package create

import (
	"log"
	"os"

	"github.com/fgazat/trc/config"
	"github.com/fgazat/trc/internal/cli"
	"github.com/fgazat/trc/internal/client"
	"github.com/spf13/cobra"
)

func Create(cfg *config.Config, creator client.Creator) *cobra.Command {
	var queue, summary, description, assignee string
	var followers []string
	cmd := &cobra.Command{
		Use:              "create",
		Short:            "Create issue",
		TraverseChildren: true,
		Aliases:          []string{"c"},
		Example:          `  trc c -q "TEST" -m "hello" -d "world"`,
		Run: func(cmd *cobra.Command, args []string) {
			createArgs := client.CreateArgs{
				Queue:       queue,
				Summary:     summary,
				Description: description,
				Assignee:    assignee,
				Followers:   followers,
			}
			log.Println(cli.StringKeyVals("Issue params", &createArgs))
			if !cfg.Force {
				ok := cli.Confirm("Create Issue?")
				if !ok {
					log.Println("Ok then")
					os.Exit(0)
				}
			}
			key, err := creator.CreateIssue(&createArgs)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf(cfg.WebBaseURL + "/" + key)
		},
	}
	cmd.Flags().StringVarP(&queue, "queue", "q", cfg.Issues.DefaultQueue, "Queue")
	cmd.Flags().StringVarP(&summary, "summary", "s", "", "Summary")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "Description of issue")
	cmd.Flags().StringVarP(&assignee, "assignee", "a", cfg.Issues.Assignee, "Assignee")
	cmd.Flags().StringSliceVar(&followers, "followers", []string{}, `Slice of followers: -fol="v1,v2" --fol="v3". Followers will be set with value: [v1 v2 v3].`)
	return cmd
}
