package pr

import (
	"encoding/json"
	"log"
	"os"

	"github.com/cli/go-gh"
)

func GetPRs() PullRequests {
	commandArgs := []string{"pr", "list", "-s", "all", "--json", "author,title,url,updatedAt,reviews"}
	out := execCommand(commandArgs)
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
