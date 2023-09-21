package experiments

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"log/slog"
	"time"
)

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
		count:  count,
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
		time.Sleep(5 * time.Second)
		branch := fmt.Sprintf("fb-%d-%v", i, GenerateUuid()[:6])
		slog.Info("Starting flow", "branch", branch)
		branchName, err := createBranch(ctx, s.client, s.Repo, branch)
		if err != nil {
			return st, err
		}
		time.Sleep(5 * time.Second)
		//slog.Info("Branch created", "name", branchName)

		_, err = createFile(ctx, s.client, s.Repo, branchName, i)
		if err != nil {
			return st, err
		}
		//slog.Info("File created", "name", file)

		time.Sleep(5 * time.Second)
		_, err = createPR(ctx, s.client, s.Repo, branchName, i)
		if err != nil {
			return st, err
		}
		//slog.Info("PR created", "name", pr)
	}
	return st, nil

}
