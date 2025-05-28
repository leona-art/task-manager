package issue

type IssueStatus interface {
	Status() string
	Cause() (value string, ok bool)
	Resolution() (value string, ok bool)
	Candidate() map[string]func() IssueStatus
}

const (
	Pending     = "pending"
	Researching = "researching"
	Resolving   = "resolving"
	Resolved    = "resolved"
)
