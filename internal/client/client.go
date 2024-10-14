package client

type CreateArgs struct {
	Queue       string
	Summary     string
	Description string
	Assignee    string
}

type Creator interface {
	CreateIssue(args *CreateArgs) (string, error)
}
