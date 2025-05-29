package issue

type IssueStatus string

const (
	Pending     IssueStatus = "pending"
	Researching IssueStatus = "researching"
	Resolving   IssueStatus = "resolving"
	Resolved    IssueStatus = "resolved"
)
