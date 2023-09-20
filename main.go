package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/google/go-github/v35/github"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func main() {
	accessToken := ""
	owner := "wings-software"
	repoName := "meenasync"
	sourceCommit := "e98fe0073c358f446d291959c14ec0135551eaad"
	sourceBranch := "create-pr-2"

	ctx := context.Background()

	client := github.NewClient(
		oauth2.NewClient(ctx, oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken},
		)),
	)

	rand.Seed(time.Now().UnixNano())

	change := "This is a simple change."
	maxRetries := 11
	maxJitter := 100

	minDelay := 3
	maxDelay := 15

	fmt.Printf("Process starting at %s\n", time.Now())
	for i := 1; i <= 50; i++ {
		branchName := fmt.Sprintf("feature-branch-%v", uuid.New())
		commitMessage := fmt.Sprintf("Commit %d", i)

		// branch
		_, _, err := client.Git.CreateRef(ctx, owner, repoName, &github.Reference{
			Ref: github.String("refs/heads/" + branchName),
			Object: &github.GitObject{
				SHA: github.String(sourceCommit),
			},
		})
		if err != nil {
			log.Fatalf("Error creating branch: %v\nResponse: %s", err, err.Error())
		}

		change = fmt.Sprintf("%s %d", change, i)

		//initialDelay := rand.Intn(151) + 30
		//fmt.Printf("Initial delay: %d seconds\n", initialDelay)
		//time.Sleep(time.Duration(initialDelay) * time.Second)

		_, _, err = client.Repositories.CreateFile(ctx, owner, repoName, "path/to/file.txt", &github.RepositoryContentFileOptions{
			Branch:  &branchName,
			Message: &commitMessage,
			Content: []byte(change),
		})
		if err != nil {
			log.Fatalf("Error creating file: %v", err)
		}

		// pull request
		prTitle := fmt.Sprintf("Pull Request %d", i)
		prBody := fmt.Sprintf("This is Pull Request %d.", i)
		//pr, resp, err := client.PullRequests.Create(ctx, owner, repoName, &github.NewPullRequest{
		//	Title: &prTitle,
		//	Body:  &prBody,
		//	Base:  &sourceBranch,
		//	Head:  &branchName,
		//})
		//if err != nil {
		//	fmt.Printf("Response Headers: %v\n", resp.Header)
		//	log.Fatalf("Error creating pull request: %v", err)
		//}
		//
		//fmt.Printf("Created pull request %d at %s: %s\n", i, time.Now(), pr.GetHTMLURL())
		//initialDelay = rand.Intn(151) + 50

		initialDelay := rand.Intn(maxDelay-minDelay+1) + minDelay
		fmt.Printf("Initial delay: %d seconds\n", initialDelay)
		time.Sleep(time.Duration(initialDelay) * time.Second)

		for retry := 7; retry <= maxRetries; retry++ {

			pr, resp, err := client.PullRequests.Create(ctx, owner, repoName, &github.NewPullRequest{
				Title: &prTitle,
				Body:  &prBody,
				Base:  &sourceBranch,
				Head:  &branchName,
			})
			if err == nil {
				fmt.Printf("Created pull request %d at %s: %s\n", i, time.Now(), pr.GetHTMLURL())
				break
			}

			fmt.Printf("Error creating pull request (Retry %d): %v\n", retry, err)
			fmt.Printf("Response Headers: %v\n", resp.Header)
			jitter := rand.Intn(maxJitter)
			if retry != maxRetries {
				//delay := time.Duration(20*(1<<(retry-1))) * time.Second
				//backoffDelay := int(math.Min(1000, float64(2*(int(math.Pow(2, float64(retry)))))))
				//backoffDelay := rand.Intn(81) + 120
				backoffDelay := int(math.Min(1024, math.Pow(2, float64(retry)))) + jitter
				fmt.Printf("Retry %d: Backoff delay: %d seconds\n", retry, backoffDelay)
				time.Sleep(time.Duration(backoffDelay) * time.Second)
			} else {
				fmt.Printf("Process stopped at %s\n", time.Now())
				log.Fatalf("Max retries reached, skipping...\n")
			}
		}
	}
	fmt.Printf("Process ended at %s\n", time.Now())
}
