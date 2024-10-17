package list

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/tabwriter"

	"github.com/fgazat/tracker/entities"
	"github.com/fgazat/trc/config"
	"github.com/fgazat/trc/internal/client"
	"github.com/fgazat/trc/internal/terminal"
	"github.com/spf13/cobra"
)

func List(cfg *config.Config, lister client.Lister) *cobra.Command {
	var open bool
	var filterIndex int
	cmd := &cobra.Command{
		Use:              "list",
		Short:            "List issues",
		TraverseChildren: true,
		Aliases:          []string{"l"},
		Example:          `  trc l -n 2 -o`,
		Run: func(cmd *cobra.Command, args []string) {
			formSelector := []string{}
			for i, filter := range cfg.Filters {
				formSelector = append(formSelector, fmt.Sprintf("%d. %s", i+1, filter.Name))
			}
			if filterIndex == 0 {
				filterIndex = terminal.SelectIndexAnswer("Select query:", formSelector, cfg.Terminal.ResultsTableHeight) + 1
			} else {
				log.Println("Selected filter: " + cfg.Filters[filterIndex-1].Name)
			}
			filterIndex = filterIndex - 1
			issues, err := lister.GetIssuesByQuery(cfg.Filters[filterIndex].Query)
			if err != nil {
				log.Fatal(err)
			}
			result := getResultMessage(issues, cfg.Terminal.SummaryMaxLength)
			if open {
				lines := strings.Split(result, "\n")
				ind := terminal.SelectIndexAnswer("Select issue:", lines, cfg.Terminal.ResultsTableHeight)
				terminal.OpenURL(fmt.Sprintf("%s/%s", cfg.WebBaseURL, issues[ind].Key))
			} else {
				log.Println(result)
			}
		},
	}
	cmd.Flags().BoolVarP(&open, "open", "o", false, "Select issue and open in browser")
	cmd.Flags().IntVarP(&filterIndex, "number", "n", 0, "Filter number")
	return cmd
}

func getResultMessage(issues []entities.Issue, summaryLenth int) string {
	buf := bytes.Buffer{}
	writer := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', tabwriter.TabIndent)
	for i, issue := range issues {
		summary := terminal.ShortenString(issue.Summary, summaryLenth)
		fmt.Fprintf(writer, "%d.\t%s\t%s\t%s\n", i+1, issue.Key, issue.Status.Display, summary)
	}
	_ = writer.Flush()
	result := strings.TrimSpace(buf.String())

	return result
}
