package experiments

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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

	fmt.Printf("Ref= %v\n", ref)
	fmt.Printf("Resp= %v\n", resp)
	return *ref.Ref, err
}

func createFile(ctx context.Context, client *github.Client, repo *Repo, branchName string, i int) (string, error) {
	commitMessage := fmt.Sprintf("Commit %d", i)
	change := fmt.Sprintf("This is a simple change %d", i)

	//initialDelay := rand.Intn(151) + 30
	//fmt.Printf("Initial delay: %d seconds\n", initialDelay)
	//time.Sleep(time.Duration(initialDelay) * time.Second)

	ref, resp, err := client.Repositories.CreateFile(ctx, repo.Owner, repo.Name, "file.txt", &github.RepositoryContentFileOptions{
		Branch:  &branchName,
		Message: &commitMessage,
		Content: []byte(change),
	})

	fmt.Printf("Ref= %v\n", ref)
	fmt.Printf("Resp= %v\n", resp)
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
	fmt.Printf("PR= %v\n", pr)
	fmt.Printf("Resp= %v\n", resp)
	return pr.GetHTMLURL(), err
}
