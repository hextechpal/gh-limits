package main

import (
	"context"
	"flag"
	"github.com/harness/gh-limits/experiments"
	"log/slog"
)

const pat = ""
const pat1 = ""

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
		stats, err := experiments.NewSerial(*token, repo, 50).Run(ctx)
		if err != nil {
			slog.Error("Error while running experiment", "err", err)
			return
		}
		slog.Info("experiment successful", "stats", stats)
	} else {
		slog.Info("Invalid message. Please specify experiment name")
	}
}
