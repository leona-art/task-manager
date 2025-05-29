package issue

type IssueRepository interface {
	Save(issue Issue) error
	GetByID(id IssueId) (value Issue, ok bool, err error)
	Update(issue Issue) error
	Delete(id IssueId) (ok bool, err error)
	List() (issues []Issue, err error)
}
