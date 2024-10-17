package create

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fgazat/trc/config"
	"github.com/fgazat/trc/internal/client"
	"github.com/fgazat/trc/internal/terminal"
	"github.com/spf13/cobra"
)

func Create(cfg *config.Config, creator client.Creator) *cobra.Command {
	var queue, summary, description, assignee string
	var followers, tags []string
	cmd := &cobra.Command{
		Use:              "create",
		Short:            "Create issue",
		TraverseChildren: true,
		Aliases:          []string{"c"},
		Example:          `  trc c -q "TEST" -s "hello" -d "world"`,
		Run: func(cmd *cobra.Command, args []string) {
			createArgs := client.CreateArgs{
				Queue:       queue,
				Summary:     summary,
				Description: description,
				Assignee:    assignee,
				Followers:   followers,
				Tags:        tags,
			}
			if err := validate(&createArgs); err != nil {
				log.Fatalf("Required fields empty: %v", err)
			}
			log.Println(terminal.StringKeyVals("Issue params", &createArgs))
			if !cfg.Force {
				ok := terminal.Confirm("Create Issue?")
				if !ok {
					log.Println("ok then...")
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
	cmd.Flags().StringSliceVar(&followers, "followers", []string{}, `Slice of followers, specifiend in one value separated with "," or in several flags.`)
	cmd.Flags().StringSliceVar(&tags, "tags", []string{}, `Slice of tags, specifiend in one value separated with "," or in several flags.`)
	return cmd
}

func validate(args *client.CreateArgs) error {
	var errs error
	if args.Queue == "" {
		errs = errors.Join(errs, fmt.Errorf("Queue"))
	}
	if args.Summary == "" {
		errs = errors.Join(errs, fmt.Errorf("Summary"))
	}
	return errs
}
