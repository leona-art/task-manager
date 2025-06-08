package issue

type IssueStatus string

const (
	Open          IssueStatus = "open"
	Investigating IssueStatus = "investigating"
	Resolving     IssueStatus = "resolving"
	Closed        IssueStatus = "closed"
)
