package pr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

var (
	ARGS_PULL_REQUESTS    = []string{"pr", "list", "-A", "@me", "-s", "all", "--json", "title,url,updatedAt"}
	ARGS_REQUESTED_REVIEW = []string{"pr", "list", "-S", "user-review-requested:@me", "-s", "all", "--json", "title,url,updatedAt"}
)

func GetPRs() PullRequests {
	out := execCommand(ARGS_PULL_REQUESTS)
	pullRequests := unmarshalPRs(out)
	return pullRequests
}

func GetReviews() PullRequests {
	out := execCommand(ARGS_REQUESTED_REVIEW)
	pullRequests := unmarshalPRs(out)
	return pullRequests
}

func execCommand(args []string) []byte {
	cmd := exec.Command("gh", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}

	return stdout.Bytes()
}

func unmarshalPRs(data []byte) PullRequests {
	var prs []PullRequest
	err := json.Unmarshal(data, &prs)
	if err != nil {
		log.Fatal(err)
	}

	return prs
}
