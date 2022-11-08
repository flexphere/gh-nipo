package pr

import (
	"encoding/json"
	"log"
	"os"

	"github.com/cli/go-gh"
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
	stdOut, stdErr, err := gh.Exec(args...)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	if stdErr.Len() > 0 {
		log.Fatalln(string(stdErr.Bytes()))
		os.Exit(1)
	}
	return stdOut.Bytes()
}

func unmarshalPRs(data []byte) PullRequests {
	var prs []PullRequest
	err := json.Unmarshal(data, &prs)
	if err != nil {
		log.Fatal(err)
	}
	return prs
}
