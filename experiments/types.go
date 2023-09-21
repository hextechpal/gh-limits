package experiments

type Stats struct {
	Name       string // Name of the experiment
	Iterations int
	Start      int64
	End        int64

	Retry map[int]int
}

type Repo struct {
	Owner        string
	Name         string
	SourceBranch string
	SourceCommit string
}
