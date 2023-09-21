package main

import (
	"context"
	"flag"
	"github.com/harness/gh-limits/experiments"
	"log/slog"
)

const pat = "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"
const pat1 = "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"

func main() {
	// Define a string flag named "message" with a default value of "Hello".
	exp := flag.String("exp", "exp1", "Specify an experiment to run")
	token := flag.String("token", pat, "Github token to be used")
	flag.Parse()

	// Check the value of the "message" flag and print accordingly.
	if *exp == "exp1" {
		ctx := context.Background()
		repo := experiments.Repo{
			Owner:        "hextechpal",
			Name:         "gh-limits",
			SourceBranch: "develop",
			SourceCommit: "ec08b69d57beedfa05a39b460b7fe85f1cd27362",
		}
		stats, err := experiments.NewSerial(*token, repo, 100).Run(ctx)
		if err != nil {
			slog.Error("Error while running experiment", "err", err)
			return
		}
		slog.Info("experiment successful", "stats", stats)
	} else if *exp == "exp2" {
		ctx := context.Background()
		repo := experiments.Repo{
			Owner:        "hextechpal",
			Name:         "gh-limits",
			SourceBranch: "develop",
			SourceCommit: "ec08b69d57beedfa05a39b460b7fe85f1cd27362",
		}
		stats, err := experiments.NewMergePulls(*token, repo).Run(ctx)
		if err != nil {
			slog.Error("Error while running experiment", "err", err)
			return
		}
		slog.Info("experiment successful", "stats", stats)
	} else {
		slog.Info("Invalid message. Please specify experiment name")
	}
}
