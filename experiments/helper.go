package experiments

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"log/slog"
	"time"
)

func createGHClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	c := oauth2.NewClient(ctx, ts)
	return github.NewClient(c)
}

func createBranch(ctx context.Context, client *github.Client, repo *Repo, branch string) (string, error) {
	ref, resp, err := client.Git.CreateRef(ctx, repo.Owner, repo.Name, &github.Reference{
		Ref: github.String("refs/heads/" + branch),
		Object: &github.GitObject{
			SHA: github.String(repo.SourceCommit),
		},
	})

	if err != nil {
		return "", err
	}

	slog.Info("Branch Created", "branch", branch, "rate", resp.Rate)
	return *ref.Ref, err
}

func createFile(ctx context.Context, client *github.Client, repo *Repo, branchName string, i int) (string, error) {
	commitMessage := fmt.Sprintf("Commit %d", i)
	change := fmt.Sprintf("This is a simple change %d", i)
	fp := fmt.Sprintf("file_%d.txt", i)

	//initialDelay := rand.Intn(151) + 30
	//fmt.Printf("Initial delay: %d seconds\n", initialDelay)
	//time.Sleep(time.Duration(initialDelay) * time.Second)

	ref, resp, err := client.Repositories.CreateFile(ctx, repo.Owner, repo.Name, fp, &github.RepositoryContentFileOptions{
		Branch:  &branchName,
		Message: &commitMessage,
		Content: []byte(change),
	})

	if err != nil {
		return "", err
	}

	slog.Info("File created", "iter", i, "Rate", resp.Rate)
	return *ref.SHA, err
}

func createPR(ctx context.Context, client *github.Client, repo *Repo, branchName string, i int) (string, error) {
	prTitle := fmt.Sprintf("Pull Request %d", i)
	prBody := fmt.Sprintf("This is Pull Request %d.", i)
	pr, resp, err := client.PullRequests.Create(ctx, repo.Owner, repo.Name, &github.NewPullRequest{
		Title: &prTitle,
		Body:  &prBody,
		Base:  &repo.SourceBranch,
		Head:  &branchName,
	})

	if err != nil {
		slog.Info("PR Failed", "iter", i, "Rate", resp.Rate)
		return "", err
	}

	//fmt.Printf("PR= %v\n", pr)
	slog.Info("PR created", "iter", i, "Rate", resp.Rate)
	return pr.GetHTMLURL(), err
}

func listPR(ctx context.Context, client *github.Client, repo *Repo) ([]int, error) {
	list, resp, err := client.PullRequests.List(ctx, repo.Owner, repo.Name, &github.PullRequestListOptions{
		State: "open",
		Base:  "develop",
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	})

	if err != nil {
		slog.Info("List PR Failed", "Rate", resp.Rate)
		return []int{}, err
	}

	prs := make([]int, len(list))
	for i, pr := range list {
		prs[i] = *pr.Number
	}
	return prs, nil

}

func mergePr(ctx context.Context, client *github.Client, repo *Repo, prs []int) ([]bool, error) {
	ans := make([]bool, len(prs))
	for i := 0; i < len(prs); i++ {
		merge, resp, err := client.PullRequests.Merge(ctx, repo.Owner, repo.Name, prs[i], fmt.Sprintf("Merging PR: %d", i), &github.PullRequestOptions{
			MergeMethod: "squash",
		})
		if err != nil {
			slog.Info("Merge PR Failed", "Rate", resp.Rate)
			return ans, err
		}
		slog.Info("Pr merged ", "pr", prs[i], "rate", resp.Rate)
		ans[i] = *merge.Merged
		time.Sleep(5 * time.Second)
	}

	return ans, nil

}

func GenerateUuid() string {
	id := uuid.New()
	return base64.RawURLEncoding.EncodeToString(id[:])
}
