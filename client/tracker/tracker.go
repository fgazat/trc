package tracker

import (
	"context"

	"github.com/fgazat/tracker"
	"github.com/fgazat/tracker/entities"
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
	)
	return &Client{
		client: client,
	}
}

type Issue struct {
	Key         string `json:"key,omitempty"`
	Queue       string `json:"queue,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateIssue(queue, summary, description string) (string, error) {
	issue := entities.Issue{
		Queue:       &entities.Entity{Key: queue},
		Summary:     summary,
		Description: description,
	}
	err := c.client.CreateIsueeAny(context.Background(), issue, &issue)
	return issue.Key, err
}
