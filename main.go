package main

import (
	"fmt"

	"github.com/flexphere/gh-nipo/pr"
)

func main() {
	createdPRs := pr.GetPRs().GetRecents().Select("Your PRs")
	reviewPRs := pr.GetReviews().GetRecents().Select("Your Reviews")

	fmt.Println(":クラッカー:")
	for _, pr := range createdPRs {
		fmt.Printf("[%s](%s)\n", pr.Title, pr.URL)
	}
	for _, pr := range reviewPRs {
		fmt.Printf("[%s](%s)\n", pr.Title, pr.URL)
	}
}
