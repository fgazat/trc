package client

type CreateArgs struct {
	Queue       string
	Summary     string
	Description string
	Assignee    string
	Followers   []string
	Tags        []string
}

type Creator interface {
	CreateIssue(args *CreateArgs) (string, error)
}
