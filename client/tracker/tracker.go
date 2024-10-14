package tracker

import (
	"context"

	"github.com/fgazat/tracker"
	"github.com/fgazat/tracker/entities"
	"github.com/fgazat/trc/client"
	"github.com/fgazat/trc/config"
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

func (c *Client) CreateIssue(args client.CreateArgs) (string, error) {
	issue := entities.Issue{
		Queue:       &entities.Entity{Key: args.Queue},
		Summary:     args.Summary,
		Description: args.Description,
	}
	err := c.client.CreateIsueeAny(context.Background(), issue, &issue)
	return issue.Key, err
}
