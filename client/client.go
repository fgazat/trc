package client

type CreateArgs struct {
	Queue       string
	Summary     string
	Description string
}

type Creator interface {
	CreateIssue(args CreateArgs) (string, error)
}
