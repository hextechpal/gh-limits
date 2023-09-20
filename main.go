package main

import (
	"flag"
	"fmt"
	"github.com/harness/gh-limits/experiments"
)

const pat = "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"

func main() {
	// Define a string flag named "message" with a default value of "Hello".
	exp := flag.String("exp", "exp1", "Specify an experiment to run")
	token := flag.String("token", pat, "Github token to be used")
	flag.Parse()

	// Check the value of the "message" flag and print accordingly.
	if *exp == "exp1" {
		repo := experiments.Repo{
			Owner:        "hextechpal",
			Name:         "gh-limits",
			SourceBranch: "develop",
			SourceCommit: "",
		}
		experiments.NewSerial(*token, repo, 1)
	} else {
		fmt.Println("Invalid message. Please specify 'Hello' or 'World'.")
	}
}
