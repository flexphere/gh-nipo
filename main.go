package main

import (
	"fmt"

	"github.com/flexphere/gh-nipo/pr"
)

func main() {
	createdPRs := pr.GetPRs().GetRecents().Select("Your PRs")
	reviewPRs := pr.GetReviews().GetRecents().Select("Your Reviews")

	if len(createdPRs)+len(reviewPRs) == 0 {
		return
	}

	fmt.Println(":クラッカー:")
	for _, pr := range append(createdPRs, reviewPRs...) {
		fmt.Printf("[%s](%s)\n", pr.Title, pr.URL)
	}
}
