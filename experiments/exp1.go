package experiments

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"log/slog"
	"time"
)

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

// Serial : This is a serial experiment that runs a request of creating and merging PRs
// with a given time in between
type Serial struct {
	client *github.Client
	Repo   *Repo

	count int
}

func NewSerial(token string, repo Repo, count int) *Serial {
	ctx := context.Background()
	c := createGHClient(ctx, token)
	return &Serial{
		client: c,
		Repo:   &repo,
	}
}

func (s *Serial) Run(ctx context.Context) (*Stats, error) {
	st := &Stats{
		Name:       "serial",
		Iterations: s.count,
		Start:      time.Now().UnixMilli(),
		Retry:      make(map[int]int),
	}
	for i := 0; i < s.count; i++ {
		st.Retry[i]++
		branchName, err := createBranch(ctx, s.client, s.Repo, fmt.Sprintf("fb-%d", i))
		if err != nil {
			return st, err
		}
		slog.Info("Branch created", "name", branchName)

		file, err := createFile(ctx, s.client, s.Repo, branchName, i)
		if err != nil {
			return st, err
		}
		slog.Info("Branch created", "name", file)

		pr, err := createPR(ctx, s.client, s.Repo, branchName, i)
		if err != nil {
			return st, err
		}
		slog.Info("PR created", "name", pr)
	}
	return st, nil

}
