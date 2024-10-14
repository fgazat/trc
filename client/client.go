package client

type Creator interface {
	CreateIssue(queue, summary, description string) (string, error)
}
