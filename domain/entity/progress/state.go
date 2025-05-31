package progress

type TransitionMap map[ProgressStatus]func() ProgressState

type ProgressState interface {
	Candidates() TransitionMap
	Solution() (value string, ok bool)
	Status() ProgressStatus
}
