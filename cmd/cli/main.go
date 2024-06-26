package main

import (
	"codebleu/internal/app/cli"
	"os"
)

func main() {
	pullRequestId := os.Getenv("BITBUCKET_PR_ID")
	if pullRequestId == "" {
		panic("Failed Get PULL REQUEST ID")
	}
	cli.Run(pullRequestId)
}
