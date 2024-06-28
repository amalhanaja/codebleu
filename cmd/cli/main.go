package main

import (
	"codebleu/internal/app/cli"
)

func main() {
	cli.Run()
	// pullRequestId := os.Getenv("BITBUCKET_PR_ID")
	// if pullRequestId == "" {
	// 	panic("Failed Get PULL REQUEST ID")
	// }
	// cli.RunE(pullRequestId)
}
