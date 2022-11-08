package main

import (
	"fmt"

	"github.com/cli/go-gh"
)

func main() {
	args := []string{"pr", "list", "-A", "@me", "-s", "all", "--json", "title,url,updatedAt"}
	stdOut, stdErr, err := gh.Exec(args...)
	if err != nil {
		fmt.Println(err)
		return
	}

	if stdErr.Len() > 0 {
		fmt.Println(stdErr)
	}

	fmt.Println(stdOut)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
