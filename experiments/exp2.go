package experiments

import (
	"context"
	"github.com/google/go-github/github"
	"log/slog"
	"time"
)

type MergePulls struct {
	client *github.Client
	Repo   *Repo
}

func NewMergePulls(token string, repo Repo) *MergePulls {
	ctx := context.Background()
	c := createGHClient(ctx, token)
	return &MergePulls{
		client: c,
		Repo:   &repo,
	}
}

func (m *MergePulls) Run(ctx context.Context) (*Stats, error) {
	stat := &Stats{
		Name:  "merge",
		Start: time.Now().UnixMilli(),
		End:   0,
		Retry: make(map[int]int),
	}
	prs, err := listPR(ctx, m.client, m.Repo)
	if err != nil {
		return stat, err
	}
	slog.Info("Found Prs", "count", len(prs))

	_, err = mergePr(ctx, m.client, m.Repo, prs)
	if err != nil {
		return stat, err

	}
	stat.End = time.Now().UnixMilli()
	return stat, nil
}
