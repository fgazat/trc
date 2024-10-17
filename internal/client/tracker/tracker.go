package tracker

import (
	"context"

	"github.com/fgazat/tracker"
	"github.com/fgazat/tracker/entities"
	"github.com/fgazat/trc/config"
	"github.com/fgazat/trc/internal/client"
)

type Client struct {
	client *tracker.Client
}

func New(cfg *config.Config) *Client {
	client := tracker.New(
		cfg.Token,
		tracker.WithBaseURL(cfg.APIBaseURL),
		tracker.WithXCloudOrgID(cfg.XCloudOrgID),
		tracker.WithXOrgID(cfg.XOrgID),
		tracker.WithUserAgent("CLI tool"),
	)
	return &Client{
		client: client,
	}
}

func (c *Client) CreateIssue(args *client.CreateArgs) (string, error) {
	issue := entities.Issue{
		Queue:       &entities.Entity{Key: args.Queue},
		Summary:     args.Summary,
		Description: args.Description,
	}
	if args.Assignee != "" {
		issue.Assignee = &entities.User{
			ID: args.Assignee,
		}
	}
	if len(args.Followers) != 0 {
		for _, u := range args.Followers {
			issue.Followers = append(issue.Followers, entities.User{ID: u})
		}
	}
	err := c.client.CreateIsueeAny(context.Background(), issue, &issue)
	return issue.Key, err
}

func (c *Client) GetIssuesByQuery(query string) ([]entities.Issue, error) {
	var issues []entities.Issue
	err := c.client.GetIssuesByQuery(context.Background(), query, &issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
}
